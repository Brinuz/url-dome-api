package randomizer

import (
	"math/rand"
	"strings"
	"url-at-minimal-api/internal/adapters/clock"
)

const (
	charSet = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"
)

// Randomizer interface
type Randomizer interface {
	RandomString(length int) string
}

// Random implements default randomizer
type Random struct {
	clock clock.Clocker
}

// New returns a valid instace of Random
func New(c clock.Clocker) *Random {
	return &Random{clock: c}
}

// RandomString generates a "random" string based on the length
func (r Random) RandomString(length int) string {
	rand.Seed(r.clock.NowUnixNano())
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	return sb.String()
}
