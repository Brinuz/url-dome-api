package router

import (
	"net/http"
	"url-at-minimal-api/internal/adapters/handlers/minify"
	"url-at-minimal-api/internal/adapters/handlers/redirect"

	"github.com/go-chi/chi"
)

// Router structure
type Router struct {
	minify   minify.Minify
	redirect redirect.Redirect
}

// New New returns a valid instace of Router
func New(m minify.Minify, r redirect.Redirect) *Router {
	return &Router{
		minify:   m,
		redirect: r,
	}
}

// Handler is the router main handler
func (r Router) Handler() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Post("/minify", r.minify.Handler)
	mux.Get("/{target}", r.redirect.Handler)

	return mux
}
