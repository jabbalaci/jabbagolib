// I/O operations.
package jio

import (
	"bufio"
	"fmt"
	"os"
)

var p = fmt.Println
var pf = fmt.Printf
var spf = fmt.Sprintf

func RawInput(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}
