package minifyurl_test

import (
	"testing"
	"url-at-minimal-api/internal/features/minifyurl"

	"github.com/stretchr/testify/assert"
)

func TestMinify(t *testing.T) {
	// Given
	mockRepo := &MockRepository{
		SaveFn: func(url, hash string) {},
	}
	mockRandom := &MockRandomizer{
		RandomStringFn: func(len int) string {
			assert.Equal(t, 7, len)
			return "AsdvReX"
		},
	}
	minifer := minifyurl.CreateMinifer(mockRepo, mockRandom)

	// When
	minifiedUrl := minifer.Minify("https://www.google.com", 7)

	// Then
	assert.Equal(t, "/AsdvReX", minifiedUrl)
	assert.Equal(t, 1, mockRepo.SaveFnCount)
}

type MockRandomizer struct {
	RandomStringFn      func(len int) string
	RandomStringFnCount int
}

func (m *MockRandomizer) RandomString(len int) string {
	m.RandomStringFnCount++
	return m.RandomStringFn(len)
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
	return mock.Find(hash)
}
