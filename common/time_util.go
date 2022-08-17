package common

import "time"

func MeasureTime(f func()) time.Duration {
	now := time.Now()
	f()
	return time.Since(now)
}
