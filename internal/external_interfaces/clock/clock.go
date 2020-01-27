package clock

// Clock is the interface created to easily manipulate clock if necessary
type Clock interface {
	NowUnixNano() int64
}
