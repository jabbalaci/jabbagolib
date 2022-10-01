// Terminal related functions.

package jconsole

import (
	"fmt"

	"golang.org/x/term"
)

// If you can't access the terminal size, call this.
// Maybe you are not even in a terminal.
// Under Windows it didn't work for me (got "not in a term"). Under Linux it worked.
func GetTerminalInfo() {
	if term.IsTerminal(0) {
		fmt.Println("in a term")
	} else {
		fmt.Println("not in a term")
	}
}

// Returns the width and the height of the terminal.
// -1 signals an error
// Under Windows it didn't work for me. Under Linux it worked.
func GetTerminalSize() (int, int) {
	if !term.IsTerminal(0) {
		// we are not in a term
		return -1, -1
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return -1, -1
	}
	return width, height
}

// Returns the width of the terminal.
// If the width is not accessible, it returns -1 .
// Under Windows it didn't work for me. Under Linux it worked.
func GetTerminalWidth() int {
	width, _ := GetTerminalSize()
	return width
}

// Returns the height of the terminal.
// If the height is not accessible, it returns -1 .
// Under Windows it didn't work for me. Under Linux it worked.
func GetTerminalHeight() int {
	_, height := GetTerminalSize()
	return height
}

// func GetTerminalWidth() int {
// 	bytes, err := exec.Command("tput", "cols").Output()
// 	s := strings.TrimSpace(string(bytes))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("#", s)
// 	n, _ := strconv.Atoi(s)
// 	return n
// }
