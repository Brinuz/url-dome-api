package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-at-minimal-api/internal/adapters/router"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Given
	mockMinifier := &MockMinify{}
	mockRedirecter := &MockRedirect{}
	router := router.New(mockMinifier, mockRedirecter, []router.Middleware{})
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Get(ms.URL + "/health-check")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestMiddlewares(t *testing.T) {
	// Given
	middlewareCalled := 0
	dummyMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middlewareCalled++
			next.ServeHTTP(w, r)
		})
	}
	mockMinifier := &MockMinify{}
	mockRedirecter := &MockRedirect{}
	middlewares := []router.Middleware{dummyMiddleware, dummyMiddleware}
	router := router.New(mockMinifier, mockRedirecter, middlewares)
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Get(ms.URL + "/health-check")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 2, middlewareCalled)
}

func TestMinifier(t *testing.T) {
	// Given
	mockMinifier := &MockMinify{
		HandlerFn: func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) },
	}
	mockRedirecter := &MockRedirect{}
	router := router.New(mockMinifier, mockRedirecter, []router.Middleware{})
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Post(ms.URL+"/minify", "application/json", nil)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestRedirecter(t *testing.T) {
	// Given
	mockMinifier := &MockMinify{}
	mockRedirecter := &MockRedirect{
		HandlerFn: func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, chi.URLParam(r, "target"), "Bcdg1A")
			w.WriteHeader(http.StatusOK)
		},
	}
	router := router.New(mockMinifier, mockRedirecter, []router.Middleware{})
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Get(ms.URL + "/Bcdg1A")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

type MockMinify struct {
	HandlerFn      func(http.ResponseWriter, *http.Request)
	HandlerFnCount int
}

func (m *MockMinify) Handler(w http.ResponseWriter, r *http.Request) {
	m.HandlerFnCount++
	m.HandlerFn(w, r)
}

type MockRedirect struct {
	HandlerFn      func(http.ResponseWriter, *http.Request)
	HandlerFnCount int
}

func (m *MockRedirect) Handler(w http.ResponseWriter, r *http.Request) {
	m.HandlerFnCount++
	m.HandlerFn(w, r)
}
