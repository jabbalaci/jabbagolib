// Console / terminal related functions.
package jconsole

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

// Reads a line from the stdin.
// It also prints a prompt if specified.
// This one doesn't return the error.
func Input(prompt string) string {
	text, _ := InputWithError(prompt)
	return text
}

// Reads a line from the stdin.
// It also prints a prompt if specified.
// This one returns the error.
func InputWithError(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return line, err
	}
	// else
	line = strings.TrimRight(line, "\r\n")
	return line, nil
}

// Reads a line from the stdin.
// It also prints a prompt if specified.
// This one uses the readline library.
func InputWithReadline(prompt string) (string, error) {
	line, err := readline.Line(prompt)
	if err != nil {
		return line, err
	}
	// else
	return line, nil
}
