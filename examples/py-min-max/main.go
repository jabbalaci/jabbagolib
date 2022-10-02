package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jpy"
)

func main() {
	values := []int{4, 20, 0, -15, -10}
	fmt.Println(values)
	fmt.Println("minimum value:", jpy.Min(values))
	fmt.Println("maximum value:", jpy.Max(values))
}
