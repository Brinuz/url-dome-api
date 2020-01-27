package minifyurl

import (
	"url-at-minimal-api/internal/external_interfaces/randomizer"
	"url-at-minimal-api/internal/external_interfaces/repository"
	"url-at-minimal-api/internal/domain"
)

// MinifyURL interface
type MinifyURL interface {
	Execute(url string, len, tries int) string
}

// Minifier is a feature used to shorten the given url
type Minifier struct {
	repository repository.Repository
	randomizer randomizer.Randomizer
}

// New returns a valid instace of Minifier
func New(rep repository.Repository, rand randomizer.Randomizer) *Minifier {
	return &Minifier{
		repository: rep,
		randomizer: rand,
	}
}

// Execute minifies the given url to a known shorter version
func (m Minifier) Execute(url string, len, tries int) string {
	if tries < 1 {
		return ""
	}
	shorten := m.randomizer.RandomString(len)
	err := m.repository.Save(url, shorten)
	if err != domain.ErrCouldNotSaveEntry {
		return shorten
	}
	return m.Execute(url, len, tries-1)
}
