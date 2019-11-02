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
	Minifier minifyurl.MinifyUrl
}

// Handler retuns an handler to be used by routing
func (m Minifier) Handler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		URL string `json:"URL"`
	}

	if r.Header.Get("Content-Type") != "application/json" || r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := m.Minifier.Minify(requestBody.URL, 7)

	respJSON, _ := json.Marshal(struct{ URL string }{r.Host + "/" + result})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(respJSON))
}

// TODO: Handle content, method, json body error
