package main

import (
	"fmt"
	"time"

	jc "github.com/jabbalaci/jabbagolib/jconsole"
)

func putChar(c byte) {
	fmt.Printf("%c", c)
	jc.MoveCursorLeft(1)
	time.Sleep(200 * time.Millisecond)
}

func main() {
	jc.ClearScreen()
	jc.HideCursor()
	jc.MoveCursorTo(2, 2)
	for i := 0; i < 4; i++ {
		putChar('*')
		jc.MoveCursorDown(1)
	}
	putChar('*')
	jc.MoveCursorRight(2)
	jc.ShowCursor()
	for i := 0; i < 4; i++ {
		putChar('*')
		jc.MoveCursorRight(2)
	}
	jc.MoveCursorLeft(2)
	jc.MoveCursorUp(1)
	for i := 0; i < 4; i++ {
		putChar('*')
		jc.MoveCursorUp(1)
	}
	jc.MoveCursorDown(1)
	jc.MoveCursorLeft(2)
	for i := 0; i < 4; i++ {
		putChar('*')
		jc.MoveCursorLeft(2)
	}

	jc.MoveCursorDown(6)
}
