package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jnum"
)

func main() {
	start := 10
	fmt.Printf("The next ten primes starting from %v:\n", start)
	fmt.Println()
	np := jnum.NextPrime(start)
	for i := 0; i < 10; i++ {
		fmt.Print(np(), ", ")
	}
	fmt.Println()
}
