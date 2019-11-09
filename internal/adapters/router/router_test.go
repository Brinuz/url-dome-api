package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-at-minimal-api/internal/adapters/handlers/minify"
	"url-at-minimal-api/internal/adapters/router"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Given
	router := router.New(minify.Minifier{})
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Get(ms.URL + "/health-check")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
