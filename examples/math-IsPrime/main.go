package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jmath"
)

func main() {
	for i := 0; i < 100; i++ {
		if jmath.IsPrime(i) {
			fmt.Printf("%v is prime\n", i)
		}
	}
}
