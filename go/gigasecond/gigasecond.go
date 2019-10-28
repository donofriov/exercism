// Package gigasecond returns the time someone's been alive for a billion seconds
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = time.Duration(1e9 * time.Second)

// AddGigasecond accepts a time (someone's birthdate) and
// returns a datetime when they will be a billion seconds old
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
