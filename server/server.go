package server

import (
	"compress/gzip"
	"context"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/guillembonet/go-templ-htmx/server/middleware"
	"github.com/guillembonet/go-templ-htmx/views"
)

type Server struct {
	server *http.Server
}

type Handler interface {
	Register(*http.ServeMux)
}

func NewServer(addr string, handler ...Handler) (*Server, error) {
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", middleware.AssetsCache(http.FileServer(views.Assets)))

	for _, h := range handler {
		h.Register(mux)
	}

	compressor := chimiddleware.NewCompressor(gzip.DefaultCompression)

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: chimiddleware.Recoverer(middleware.Logger(compressor.Handler(mux))),
		},
	}, nil
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
