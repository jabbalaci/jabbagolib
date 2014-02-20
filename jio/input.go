package jio

import (
	"bufio"
	"fmt"
	"os"
)

func RawInput(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}
