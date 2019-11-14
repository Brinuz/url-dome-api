package minify

import (
	"encoding/json"
	"net/http"
	"url-at-minimal-api/internal/features/minifyurl"
)

// Minify interface
type Minify interface {
	Handler(http.ResponseWriter, *http.Request)
}

// Minifier implments default Minifier
type Minifier struct {
	minifier minifyurl.MinifyURL
}

// New returns a valid instace of Minifier
func New(m minifyurl.MinifyURL) *Minifier {
	return &Minifier{
		minifier: m,
	}
}

// Handler retuns an handler to be used by routing
func (m Minifier) Handler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		URL string `json:"URL"`
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := m.minifier.Minify(requestBody.URL, 7)

	respJSON, _ := json.Marshal(struct{ URL string }{"/" + result})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(respJSON))
}
