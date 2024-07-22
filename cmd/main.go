package main

import (
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guillembonet/go-templ-htmx/server"
	"github.com/guillembonet/go-templ-htmx/server/handlers"
)

var (
	flagLogLevel = flag.String("log-level", slog.LevelDebug.String(), "log level")
	flagAddress  = flag.String("address", ":8080", "server address")
)

func main() {
	flag.Parse()

	if flagLogLevel == nil || *flagLogLevel == "" {
		slog.Error("log-level is required")
		os.Exit(1)
	}

	logLevel, err := parseLevel(*flagLogLevel)
	if err != nil {
		slog.Error("failed to parse log level", slog.Any("err", err), slog.String("log-level", *flagLogLevel))
		os.Exit(1)
	}

	slog.SetLogLoggerLevel(logLevel)

	if flagAddress == nil || *flagAddress == "" {
		slog.Error("address is required")
		os.Exit(1)
	}

	server, err := server.NewServer(*flagAddress, handlers.NewStatic())
	if err != nil {
		slog.Error("failed to create server", slog.Any("err", err))
		os.Exit(1)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)

		slog.Info("starting server", slog.String("address", *flagAddress))

		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", slog.Any("err", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		slog.Info("Shutting down gracefully...")
	case <-stopped:
		slog.Error("Server stopped unexpectedly, shutting down...")
	}

	if err := server.Stop(10 * time.Second); err != nil {
		slog.Error("Server failed to shutdown gracefully", slog.Any("err", err))
		os.Exit(1)
	}
}

func parseLevel(s string) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}
