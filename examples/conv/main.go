package main

import (
	"fmt"
	"strconv"

	"github.com/jabbalaci/jabbagolib/jconv"
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
	fmt.Println("---")
	pi := 3.14159
	result := jconv.Float64ToStr(pi)
	fmt.Println(result)
}
