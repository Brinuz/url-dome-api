package redirecturl

import (
	"url-at-minimal-api/internal/external_interfaces/repository"
)

// RedirectURL interface
type RedirectURL interface {
	Execute(hash string) string
}

// Redirecter is a feature to redirect to the hash's url
type Redirecter struct {
	repository repository.Repository
}

// New returns a valid instace of Minifier
func New(rep repository.Repository) *Redirecter {
	return &Redirecter{
		repository: rep,
	}
}

// Execute returns the original minified url based on the hash
func (m Redirecter) Execute(hash string) string {
	return m.repository.Find(hash)
}
