package minifyurl

import (
	"url-at-minimal-api/internal/adapters/persistence"
	"url-at-minimal-api/internal/adapters/randomizer"
)

// Minifier is a feature used to shorten the given url
type Minifier struct {
	repository persistence.Repository
	randomizer randomizer.Randomizer
}

// CreateMinifer returns a valid instace of Minifier
func CreateMinifer(rep persistence.Repository, rand randomizer.Randomizer) *Minifier {
	return &Minifier{
		repository: rep,
		randomizer: rand,
	}
}

// Minify minifies the given url to a known shorter version
func (m *Minifier) Minify(url string, len int) string {
	shorten := m.randomizer.RandomString(len)
	m.repository.Save(url, shorten)
	return "/" + shorten
}
