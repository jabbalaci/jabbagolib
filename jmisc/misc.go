// Miscellaneous stuff.
//
// If you can't decide where to put a function, place it here
// and later you move it to a proper package.
package jmisc

import (
	"fmt"
)

var p = fmt.Println
var pf = fmt.Printf
var spf = fmt.Sprintf

// custom assert function
// it panics if the expression is false
func Assert(expr bool, msg string) {
	if !expr {
		panic(spf("Assertion error: %s.", msg))
	}
}
