package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jassert"
)

func main() {
	x := -2
	// x := 2
	jassert.Assert(x >= 0, "x cannot be negative")

	fmt.Println("x = ", x)
}
