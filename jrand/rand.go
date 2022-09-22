// Package jrand facilitates work with random numbers.
package jrand

import (
	"math/rand"
	"time"

	"github.com/jabbalaci/jabbagolib/jmisc"
)

// Seed sets the random seed correctly.
// Tip: call it in init().
// The seed value is returned but it can be discarded.
func Seed() int64 {
	val := time.Now().UTC().UnixNano()
	rand.Seed(val)
	return val
}

// RandInt returns a random integer from the [lo, hi] closed interval.
func RandInt(lo, hi int) int {
	jmisc.Assert(hi >= lo, "upper limit must be >= lower limit")
	return lo + rand.Intn(hi-lo+1)
}

// ShuffleSliceInt shuffles a []int in place.
// This is an implementation of the Fisher and Yates algorithm
// as seen at http://en.wikipedia.org/wiki/Knuth_shuffle .
func ShuffleSliceInt(items []int) {
	i := len(items)
	for i > 0 {
		i--
		j := rand.Intn(i + 1) // 0 <= j <= i
		items[j], items[i] = items[i], items[j]
	}
}
