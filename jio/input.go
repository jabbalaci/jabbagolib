// Package jio facilitates I/O operations.
package jio

import (
	"bufio"
	"fmt"
	"os"
)

// Input reads a line from the stdin.
// It also prints a prompt if specified.
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}
