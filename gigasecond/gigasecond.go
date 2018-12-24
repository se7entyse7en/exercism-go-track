// Package gigasecond implements the "Gigasecond" exercise.
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond returns the given time.Time plus 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow(10, 9)))
}
