// Random numbers.
package jrand

import (
	"github.com/jabbalaci/jabbagolib/misc"
	"math/rand"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf
var spf = fmt.Sprintf

// set random seed correctly
// call it in init()
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// return a random int from the [lo, hi] closed interval
func RandInt(lo, hi int) int {
	misc.Assert(hi >= lo, "upper limit must be >= lower limit")
	return lo + rand.Intn(hi-lo+1)
}

// Fisher and Yates algorithm
// http://en.wikipedia.org/wiki/Knuth_shuffle
// shuffles in place
func ShuffleSliceInt(items []int) {
	i := len(items)
	for i > 0 {
		i--
		j := rand.Intn(i + 1) // 0 <= j <= i
		items[j], items[i] = items[i], items[j]
	}
}
