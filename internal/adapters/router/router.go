package router

import (
	"net/http"
	"url-at-minimal-api/internal/adapters/handlers/minify"

	"github.com/go-chi/chi"
)

// Router structure
type Router struct {
	minify minify.Minify
}

// New New returns a valid instace of Router
func New(m minify.Minify) *Router {
	return &Router{
		minify: m,
	}
}

// Handler is the router main handler
func (r Router) Handler() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Post("/minify", r.minify.Handler)

	return mux
}
