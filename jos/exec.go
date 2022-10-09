package jos

import (
	"log"
	"os/exec"
	"strings"

	"github.com/google/shlex"
)

// Execute an external simple command and return its output.
// The command must be simple, i.e. no pipes, no redirection.
func ExecCmdAndGetOutput(cmd string) string {
	fields, err := shlex.Split(cmd)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := exec.Command(fields[0], fields[1:]...).Output()
	s := strings.TrimSpace(string(bytes))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
