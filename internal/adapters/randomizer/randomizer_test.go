package randomizer_test

import (
	"testing"
	"url-at-minimal-api/internal/adapters/randomizer"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	// Given
	clockMock := MockClock{}
	random := randomizer.New(clockMock)

	// When
	randomString := random.RandomString(5)

	// Then
	assert.Equal(t, "gTXZN", randomString)
}

type MockClock struct {
}

func (m MockClock) NowUnixNano() int64 {
	return 1572657332 //2018/11/02
}
