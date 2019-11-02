package minifyurl

import (
	"url-at-minimal-api/internal/adapters/randomizer"
	"url-at-minimal-api/internal/adapters/repository"
)

type MinifyUrl interface {
	Minify(url string, len int) string
	Deminify(hash string) string
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

// Minify minifies the given url to a known shorter version
func (m Minifier) Minify(url string, len int) string {
	shorten := m.randomizer.RandomString(len)
	m.repository.Save(url, shorten)
	return shorten
}

// Deminify returns the original minified url based on the hash
func (m Minifier) Deminify(hash string) string {
	return m.repository.Find(hash)
}
