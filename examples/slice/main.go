package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jslice"
)

func main() {
	li := []string{"aa", "bb", "cc"}
	li, found := jslice.Remove(li, "bb")
	fmt.Println(li, found) // Output: [aa cc] true
	li, found = jslice.Remove(li, "bb")
	fmt.Println(li, found) // Output: [aa cc] false
}
