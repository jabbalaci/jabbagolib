package main

import (
	"fmt"
	"strings"

	"github.com/jabbalaci/jabbagolib/jconsole"
)

func main() {
	name := jconsole.Input("Your name: ")
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)
}
