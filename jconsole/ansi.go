// ANSI escape codes. More info here: https://en.wikipedia.org/wiki/ANSI_escape_code

package jconsole

import "fmt"

// clear entire screen
func ClearScreen() {
	fmt.Print("\x1Bc")
}

// hide cursor
// call it after ClearScreen()
func HideCursor() {
	fmt.Print("\x1b[?25l")
}

// show cursor
func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

// move the cursor to row n, column m
func MoveCursorTo(row, column int) {
	fmt.Printf("\x1b[%v;%vH", row, column)
}

// move cursor up n rows
func MoveCursorUp(n int) {
	fmt.Printf("\x1b[%vA", n)
}

// move cursor down n rows
func MoveCursorDown(n int) {
	fmt.Printf("\x1b[%vB", n)
}

// move cursor back n columns
func MoveCursorLeft(n int) {
	fmt.Printf("\x1b[%vD", n)
}

// move cursor forward n columns
func MoveCursorRight(n int) {
	fmt.Printf("\x1b[%vC", n)
}
