package redirect_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-at-minimal-api/internal/adapters/handlers/redirect"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	// Given
	mockRedictUrl := &MockRedictUrl{
		URLFn: func(hash string) string {
			assert.Equal(t, "AsdcBV1", hash)
			return "https://dummy.url"
		},
	}
	handler := redirect.New(mockRedictUrl)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("target", "AsdcBV1")
	req := httptest.NewRequest("GET", "https://mini.fy/{target}", nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rec := httptest.NewRecorder()

	// When
	http.HandlerFunc(handler.Handler).ServeHTTP(rec, req)

	// Then
	assert.Equal(t, http.StatusMovedPermanently, rec.Code)
	assert.Equal(t, "https://dummy.url", rec.HeaderMap.Get("Location"))

	assert.Equal(t, 1, mockRedictUrl.URLFnCount)
}

type MockRedictUrl struct {
	URLFn      func(hash string) string
	URLFnCount int
}

func (m *MockRedictUrl) URL(hash string) string {
	m.URLFnCount++
	return m.URLFn(hash)
}
