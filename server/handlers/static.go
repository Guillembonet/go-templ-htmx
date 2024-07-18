package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/go-templ-htmx/server"
	"github.com/guillembonet/go-templ-htmx/views/home"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

func (*Static) Root(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (*Static) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, home.Home(), "Home", "homepage"))
}

func (s *Static) Register(r *gin.RouterGroup) {
	r.GET("/", s.Root)
	r.GET("/home", s.Home)
}
