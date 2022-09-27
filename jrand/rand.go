// Package jrand facilitates work with random numbers.
package jrand

import (
	"math/rand"
	"time"

	"github.com/jabbalaci/jabbagolib/jassert"
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
	jassert.Assert(hi >= lo, "upper limit must be >= lower limit")
	return lo + rand.Intn(hi-lo+1)
}

// Returns a random integer from the [lo, hi) interval.
// `lo` is included, `hi` is excluded.
func RandRange(lo, hi int) int {
	jassert.Assert(hi > lo, "upper limit must be > lower limit")
	return lo + rand.Intn(hi-lo)
}

// Shuffles a slice in place.
func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}
