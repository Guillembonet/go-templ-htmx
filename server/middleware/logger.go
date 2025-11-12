package middleware

import (
	"log/slog"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ww := chimiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		elapsedTime := time.Since(startTime)
		path := r.URL.Path
		raw := r.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}
		statusCode := ww.Status()

		attrs := []any{
			slog.Duration("latency", elapsedTime),
			slog.String("method", r.Method),
			slog.Int("status", statusCode),
			slog.String("path", path),
		}

		msg := "request"
		if statusCode >= 500 {
			slog.Error(msg, attrs...)
			return
		}
		slog.Debug(msg, attrs...)
	})
}
