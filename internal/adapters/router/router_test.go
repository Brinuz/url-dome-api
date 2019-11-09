package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-at-minimal-api/internal/adapters/router"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Given
	mockMinifier := &MockMinify{}
	router := router.New(mockMinifier)
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Get(ms.URL + "/health-check")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestMinifier(t *testing.T) {
	// Given
	mockMinifier := &MockMinify{
		HandlerFn: func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) },
	}
	router := router.New(mockMinifier)
	ms := httptest.NewServer(router.Handler())
	defer ms.Close()

	// When
	resp, err := http.Post(ms.URL+"/minify", "application/json", nil)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

type MockMinify struct {
	HandlerFn      func(http.ResponseWriter, *http.Request)
	HandlerFnCount int
}

func (m *MockMinify) Handler(w http.ResponseWriter, r *http.Request) {
	m.HandlerFnCount++
	m.HandlerFn(w, r)
}
