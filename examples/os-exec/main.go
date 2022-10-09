package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jos"
)

func main() {
	cmd := "/bin/date"
	fmt.Println("# calling:", cmd)
	output := jos.ExecCmdAndGetOutput(cmd)
	fmt.Println("# output:", output)
}
