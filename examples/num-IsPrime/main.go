package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jnum"
)

func main() {
	for i := 0; i < 100; i++ {
		if jnum.IsPrime(i) {
			fmt.Printf("%v is prime\n", i)
		}
	}
}
