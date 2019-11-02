package minifyurl_test

import (
	"testing"
	"url-at-minimal-api/internal/features/minifyurl"

	"github.com/stretchr/testify/assert"
)

func TestMinify(t *testing.T) {
	// Given
	mockRepo := &MockRepository{
		SaveFn: func(url, hash string) {
			assert.Equal(t, "https://www.google.com", url)
		},
	}
	mockRandom := &MockRandomizer{
		RandomStringFn: func(length int) string {
			assert.Equal(t, 7, length)
			return "AsdvRe0"
		},
	}
	minifer := minifyurl.New(mockRepo, mockRandom)

	// When
	minifiedUrl := minifer.Minify("https://www.google.com", 7)

	// Then
	assert.Equal(t, "AsdvRe0", minifiedUrl)
	assert.Equal(t, 1, mockRepo.SaveFnCount)
}

func TestDeminify(t *testing.T) {
	// Given
	mockRepo := &MockRepository{
		FindFn: func(hash string) string {
			assert.Equal(t, "Casd1cV", hash)
			return "https://www.google.com"
		},
	}
	mockRandom := &MockRandomizer{}
	minifer := minifyurl.New(mockRepo, mockRandom)

	// When
	url := minifer.Deminify("Casd1cV")

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
	SaveFn      func(url, hash string)
	SaveFnCount int
	FindFn      func(hash string) string
	FindFnCount int
}

func (mock *MockRepository) Save(url, hash string) {
	mock.SaveFnCount++
	mock.SaveFn(url, hash)
}
func (mock *MockRepository) Find(hash string) string {
	mock.FindFnCount++
	return mock.FindFn(hash)
}
