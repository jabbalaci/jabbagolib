package main

import (
	"fmt"
	"strings"

	"github.com/jabbalaci/jabbagolib/jconsole"
	"github.com/jabbalaci/jabbagolib/jconsole/fancy"
)

func main() {
	fmt.Println("Normal input:")
	fmt.Println("=============")
	name := jconsole.Input("Your name: ")
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)
	fmt.Println()
	fmt.Println("Input with readline support:")
	fmt.Println("============================")
	name = fancy.Input("Your name: ")
	// fmt.Printf("%q\n", name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)
}
