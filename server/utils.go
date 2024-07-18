package server

import (
	"log/slog"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/guillembonet/go-templ-htmx/views/layouts"
)

const (
	// HXRequestHeader is the header that indicates that the request is an htmx request.
	HXRequestHeader = "HX-Request"
	// HXHistoryRestoreRequestHeader is the header that indicates that the request is a history restore request.
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
)

// requestsFullPage returns true if the request should return a full page, false if it should return a partial page.
func requestsFullPage(c *gin.Context) bool {
	htmxRequest := c.GetHeader(HXRequestHeader) == "true"
	if !htmxRequest {
		return true
	}
	restoreRequest := c.GetHeader(HXHistoryRestoreRequestHeader) == "true"
	if restoreRequest {
		slog.Debug("restoring history")
	}
	return restoreRequest
}

func WithBase(c *gin.Context, component templ.Component, title, description string) templ.Component {
	return layouts.WithBase(component, title, description, requestsFullPage(c))
}
