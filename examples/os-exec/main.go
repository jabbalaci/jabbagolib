package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jos"
)

func main() {
	if jos.IsLinux() {
		cmd := "date"
		fmt.Println("# calling:", cmd)
		output := jos.ExecCmdAndGetOutput(cmd)
		fmt.Println("# output:", output)
		fmt.Println("---")
		command := "sh -c \"for x in `seq 1 3`; do echo test; sleep 0.1; done\""
		jos.ExecCmdAndSeeOutput(command)
		fmt.Println("---")
	}

	var prg string
	if jos.IsWindows() {
		prg = `c:\Program Files\Notepad++\notepad++.exe`
	} else {
		prg = "gedit"
	}
	jos.ExecCmdInBackground(prg)
}
