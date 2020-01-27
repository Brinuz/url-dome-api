package redirect

import (
	"net/http"
	"url-at-minimal-api/internal/use_cases/redirecturl"

	"github.com/go-chi/chi"
)

// Redirect interface
type Redirect interface {
	Handler(http.ResponseWriter, *http.Request)
}

// Redirecter implments default Redirect
type Redirecter struct {
	redirecter redirecturl.RedirectURL
}

// New returns a valid instace of Redirecter
func New(r redirecturl.RedirectURL) *Redirecter {
	return &Redirecter{
		redirecter: r,
	}
}

// Handler retuns an handler to be used by routing
func (m Redirecter) Handler(w http.ResponseWriter, r *http.Request) {
	url := m.redirecter.Execute(chi.URLParam(r, "target"))
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
