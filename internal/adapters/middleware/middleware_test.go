package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-at-minimal-api/internal/adapters/middleware"

	"github.com/stretchr/testify/assert"
)

func ok(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }

func TestSecurity(t *testing.T) {
	// Given
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handler := http.HandlerFunc(ok)

	// When
	middleware.Security(handler).ServeHTTP(rec, req)

	// Then
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}
