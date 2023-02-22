package main

import (
	"fmt"
	"strings"

	"github.com/jabbalaci/jabbagolib/jconsole"
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
	name, _ = jconsole.InputWithReadline("Your name: ")
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)
}
