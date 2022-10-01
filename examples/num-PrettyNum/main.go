package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jnum"
)

func main() {
	numbers := []int{1977, 10267, 2147483647}
	for _, n := range numbers {
		fmt.Printf("%v\t->\t%v\n", n, jnum.PrettyNum(n))
	}
}
