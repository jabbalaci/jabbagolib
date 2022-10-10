package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jfile"
)

func main() {
	fname := "input.txt"
	content, _ := jfile.Read(fname)
	fmt.Printf("%q\n", content)
	fmt.Println("---")
	lines, _ := jfile.ReadLines(fname)
	for _, line := range lines {
		fmt.Printf("%q\n", line)
	}
	fmt.Println("---")
	numbers, _ := jfile.ReadLinesAsInts("numbers.txt")
	for _, number := range numbers {
		fmt.Printf("%d (%T)\n", number, number)
	}
}
