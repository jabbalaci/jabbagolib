package jpy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input reads a line from the stdin.
// It also prints a prompt if specified.
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return line, err
	}
	// else
	return strings.TrimRight(line, "\n"), nil
}
