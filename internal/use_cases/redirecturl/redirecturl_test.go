package redirecturl_test

import (
	"testing"
	"url-at-minimal-api/internal/use_cases/redirecturl"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	// Given
	mockRepo := &MockRepository{
		FindFn: func(hash string) string {
			assert.Equal(t, "Casd1cV", hash)
			return "https://www.google.com"
		},
	}
	redirect := redirecturl.New(mockRepo)

	// When
	url := redirect.Execute("Casd1cV")

	// Then
	assert.Equal(t, "https://www.google.com", url)
	assert.Equal(t, 1, mockRepo.FindFnCount)
}

type MockRandomizer struct {
	RandomStringFn      func(length int) string
	RandomStringFnCount int
}

func (m *MockRandomizer) RandomString(length int) string {
	m.RandomStringFnCount++
	return m.RandomStringFn(length)
}

type MockRepository struct {
	SaveFn      func(url, hash string) error
	SaveFnCount int
	FindFn      func(hash string) string
	FindFnCount int
}

func (mock *MockRepository) Save(url, hash string) error {
	mock.SaveFnCount++
	return mock.SaveFn(url, hash)
}
func (mock *MockRepository) Find(hash string) string {
	mock.FindFnCount++
	return mock.FindFn(hash)
}
