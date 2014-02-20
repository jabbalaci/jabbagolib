// Package jrand facilitates work with random numbers.
package jrand

import (
	"github.com/jabbalaci/jabbagolib/misc"
	"math/rand"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf
var spf = fmt.Sprintf

// Seed sets the random seed correctly.
// Tip: call it in init().
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandInt returns a random integer from the [lo, hi] closed interval.
func RandInt(lo, hi int) int {
	misc.Assert(hi >= lo, "upper limit must be >= lower limit")
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
