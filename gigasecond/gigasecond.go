// Solution to the Gigasecond exercism exercise. Calculates the point in time
// when 10^9 seconds will have elapsed starting at supplied time.
package gigasecond

import (
	"math"
	"time"
)

// Adds 10^9 seconds to supplied Time. Returns a new Time struct.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}
