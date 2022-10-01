package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jabbalaci/jabbagolib/jpy"
)

func main() {
	words := []string{"war", "cup", "water", "tree", "storm"}
	// result := jpy.Map(words, func(s string) string {
	// return strings.ToUpper(s)
	// })
	result := jpy.Map(words, strings.ToUpper)
	fmt.Println(result)

	fmt.Println("---")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	squares := jpy.Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println(squares)

	fmt.Println("---")
	as_strings := jpy.Map(numbers, func(n int) string {
		return strconv.Itoa(n)
	})
	fmt.Printf("%q", as_strings)
}
