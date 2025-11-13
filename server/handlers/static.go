package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/guillembonet/go-templ-htmx/server"
	"github.com/guillembonet/go-templ-htmx/views/error_pages"
	"github.com/guillembonet/go-templ-htmx/views/home"
)

type Static struct {
	home     templ.Component
	notFound templ.Component
}

func NewStatic() *Static {
	return &Static{
		home:     home.Home(),
		notFound: error_pages.NotFound(),
	}
}

func (s *Static) Root() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			server.WithBase(s.notFound, "Not found", "", templ.WithStatus(http.StatusNotFound)).ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	})
}

func (s *Static) Home() http.Handler {
	return server.WithBase(s.home, "Home", "homepage")
}

func (s *Static) Register(m *http.ServeMux) {
	m.Handle("GET /", s.Root())
	m.Handle("GET /home", s.Home())
}
