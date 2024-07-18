package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	if raw != "" {
		path = path + "?" + raw
	}
	statusCode := c.Writer.Status()

	attrs := []any{
		slog.Duration("latency", time.Since(start)),
		slog.String("method", c.Request.Method),
		slog.Int("status", statusCode),
		slog.String("path", path),
	}

	msg := "request"
	if statusCode >= 500 {
		slog.Error(msg, attrs...)
		return
	}
	slog.Debug(msg, attrs...)
}
