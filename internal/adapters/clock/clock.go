package clock

import "time"

// Clocker is the interface created to easily manipulate clock if necessary
type Clocker interface {
	NowUnixNano() int64
}

// Clock implements the Clocker interface
type Clock struct {
}

// NowUnixNano returns the current time as Unix time
func (c Clock) NowUnixNano() int64 {
	return time.Now().UnixNano()
}
