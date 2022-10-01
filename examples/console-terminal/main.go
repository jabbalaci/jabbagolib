package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jconsole"
)

func main() {
	jconsole.GetTerminalInfo()

	w, h := jconsole.GetTerminalSize()
	fmt.Printf("width: %v, height: %v\n", w, h)
}
