package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "12345"
	sum := 0
	for _, c := range s {
		// sum += jconv.RuneToDigit(c)
		n, _ := strconv.Atoi(string(c))
		sum += n
	}
	fmt.Println(sum)
}
