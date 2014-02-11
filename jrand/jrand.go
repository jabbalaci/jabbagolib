package jrand

import (
	"github.com/jabbalaci/jabbagolib/misc"
	"math/rand"
	"time"
)

// set random seed correctly
// call it in init()
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// return a random int from the [l, u] closed interval
func RandInt(l, u int) int {
	misc.Assert(u >= l, "upper limit must be >= lower limit")
	return l + rand.Intn(u-l+1)
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
