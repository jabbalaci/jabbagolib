package misc

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
