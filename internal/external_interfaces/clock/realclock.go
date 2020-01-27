package clock

import "time"

// RealClock implements the Clock interface
type RealClock struct {
}

// New returns a valid instace of RealClock
func New() *RealClock {
	return &RealClock{}
}

// NowUnixNano returns the current time as Unix time
func (c RealClock) NowUnixNano() int64 {
	return time.Now().UnixNano()
}
