package minifyurl_test

import (
	"testing"
	"url-at-minimal-api/internal/domain"
	"url-at-minimal-api/internal/use_cases/minifyurl"

	"github.com/stretchr/testify/assert"
)

func TestMinify(t *testing.T) {
	// Given
	mockRepo := &MockRepository{
		SaveFn: func(url, hash string) error {
			assert.Equal(t, "https://www.google.com", url)
			return nil
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
	minifiedUrl := minifer.Execute("https://www.google.com", 7, 2)

	// Then
	assert.Equal(t, "AsdvRe0", minifiedUrl)
	assert.Equal(t, 1, mockRepo.SaveFnCount)
}

func TestMinifyOnHashCollision(t *testing.T) {
	// Given
	collisions := 0

	mockRepo := &MockRepository{
		SaveFn: func(url, hash string) error {
			var err error
			switch collisions {
			case 0:
				err = domain.ErrCouldNotSaveEntry
			case 1:
				err = nil
			}
			collisions++
			return err
		},
	}
	mockRandom := &MockRandomizer{
		RandomStringFn: func(length int) string {
			var hash string
			switch collisions {
			case 0:
				hash = "Bsdf52S"
			case 1:
				hash = "AsdvRe0"
			}
			return hash
		},
	}
	minifer := minifyurl.New(mockRepo, mockRandom)

	// When
	minifiedUrl := minifer.Execute("https://www.google.com", 7, 2)

	// Then
	assert.Equal(t, "AsdvRe0", minifiedUrl)
	assert.Equal(t, 2, mockRepo.SaveFnCount)
	assert.Equal(t, 2, mockRandom.RandomStringFnCount)
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
