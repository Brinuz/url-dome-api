package rest

import (
	"net/http"
	"url-at-minimal-api/internal/external_interfaces/handlers/minify"
	"url-at-minimal-api/internal/external_interfaces/handlers/redirect"

	"github.com/go-chi/chi"
)

// Middleware represents a middleware handler
type Middleware func(next http.Handler) http.Handler

// Rest structure
type Rest struct {
	minify      minify.Minify
	redirect    redirect.Redirect
	middlewares []Middleware
}

// New New returns a valid instace of Rest
func New(m minify.Minify, r redirect.Redirect, mw []Middleware) *Rest {
	return &Rest{
		minify:      m,
		redirect:    r,
		middlewares: mw,
	}
}

// Handler is the router main handler
func (r Rest) Handler() *chi.Mux {
	mux := chi.NewRouter()

	for _, mw := range r.middlewares {
		mux.Use(mw)
	}

	mux.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Post("/minify", r.minify.Handler)
	mux.Get("/{target}", r.redirect.Handler)

	return mux
}
