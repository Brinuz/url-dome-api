package minify_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-at-minimal-api/internal/adapters/handlers/minify"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	// Given
	mockMinifyUrl := &MockMinifyUrl{
		MinifyFn: func(url string, len int) string {
			assert.Equal(t, 7, len)
			assert.Equal(t, "https://dummy.url", url)
			return "AsdcBV1"
		},
	}
	handler := minify.New(mockMinifyUrl)
	req := httptest.NewRequest("POST", "https://mini.fy/", strings.NewReader(`{"URL":"https://dummy.url"}`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// When
	http.HandlerFunc(handler.Handler).ServeHTTP(rec, req)
	result := rec.Result()

	// Then
	body, _ := ioutil.ReadAll(result.Body)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json")
	assert.JSONEq(t, `{"URL":"mini.fy/AsdcBV1"}`, string(body))
	assert.Equal(t, 1, mockMinifyUrl.MinifyFnCount)
}

func TestHandlerInvalidContent(t *testing.T) {
	// Given
	mockMinifyUrl := &MockMinifyUrl{}
	handler := minify.New(mockMinifyUrl)
	req := httptest.NewRequest("POST", "/", strings.NewReader("{invalid_json}"))

	rec := httptest.NewRecorder()

	// When
	http.HandlerFunc(handler.Handler).ServeHTTP(rec, req)

	// Then
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, 0, mockMinifyUrl.MinifyFnCount)
}

func TestHandlerBadJSON(t *testing.T) {
	// Given
	mockMinifyUrl := &MockMinifyUrl{}
	handler := minify.New(mockMinifyUrl)
	req := httptest.NewRequest("POST", "/", strings.NewReader(`1234`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// When
	http.HandlerFunc(handler.Handler).ServeHTTP(rec, req)

	// Then
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, 0, mockMinifyUrl.MinifyFnCount)
}

type MockMinifyUrl struct {
	MinifyFn      func(url string, len int) string
	MinifyFnCount int
}

func (m *MockMinifyUrl) Minify(url string, len int) string {
	m.MinifyFnCount++
	return m.MinifyFn(url, len)
}
