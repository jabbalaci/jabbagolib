package main

import (
	"fmt"
	"strconv"

	"github.com/jabbalaci/jabbagolib/jconsole"
	"github.com/jabbalaci/jabbagolib/jnum"
)

func main() {
	fmt.Println("Provide two integers")
	s1 := jconsole.Input("Number 1: ")
	n1, _ := strconv.Atoi(s1)
	s2 := jconsole.Input("Number 2: ")
	n2, _ := strconv.Atoi(s2)
	result := jnum.Pow(int64(n1), int64(n2))
	fmt.Printf("%v^%v = %v\n", n1, n2, result)
}
