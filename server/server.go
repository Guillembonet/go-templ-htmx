package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/guillembonet/go-templ-htmx/server/middleware"
	"github.com/guillembonet/go-templ-htmx/views/assets"
	"github.com/guillembonet/go-templ-htmx/views/error_pages"
)

type Server struct {
	server *http.Server
}

type Handler interface {
	Register(*gin.RouterGroup)
}

func NewServer(addr string, handler ...Handler) (*Server, error) {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()
	g.Use(middleware.Logger, gin.Recovery(),
		middleware.AssetsCache, gzip.Gzip(gzip.DefaultCompression))
	g.HTMLRender = &templRenderer{}

	g.StaticFS("/assets", http.FS(assets.Assets))

	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "", WithBase(c, error_pages.NotFound(), "Not found", ""))
	})

	rg := g.Group("/")
	for _, h := range handler {
		h.Register(rg)
	}

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: g,
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
