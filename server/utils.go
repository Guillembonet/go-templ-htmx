package server

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/guillembonet/go-templ-htmx/views/layouts"
)

const (
	// HXRequestHeader is the header that indicates that the request is an htmx request.
	HXRequestHeader = "HX-Request"
	// HXHistoryRestoreRequestHeader is the header that indicates that the request is a history restore request.
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
)

// requestsFullPage returns true if the request should return a full page, false if it should return a partial page.
func requestsFullPage(r *http.Request) bool {
	htmxRequest := r.Header.Get(HXRequestHeader) == "true"
	if !htmxRequest {
		return true
	}
	restoreRequest := r.Header.Get(HXHistoryRestoreRequestHeader) == "true"
	if restoreRequest {
		slog.Debug("history restore request", slog.String("path", r.URL.Path))
	}
	return restoreRequest
}

func WithBase(component templ.Component, title, description string, options ...func(*templ.ComponentHandler)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(layouts.WithBase(component, title, description, requestsFullPage(r)), options...).ServeHTTP(w, r)
	})
}
