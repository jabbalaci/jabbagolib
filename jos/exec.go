package jos

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/shlex"
)

// good post: https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

// Execute an external simple command and return its output.
func ExecCmdAndGetOutput(cmd string) string {
	fields, err := shlex.Split(cmd)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := exec.Command(fields[0], fields[1:]...).Output() // waits for it to complete
	s := strings.TrimSpace(string(bytes))
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// Execute an external simple command and see its output.
// Similar to Python's os.system(...) call.
func ExecCmdAndSeeOutput(command string) {
	fields, err := shlex.Split(command)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(fields[0], fields[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run() // waits for it to complete
	if err != nil {
		log.Fatal(err)
	}
}

// Execute an external command but don't see its output.
func ExecCmdWithoutOutput(command string) {
	fields, err := shlex.Split(command)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(fields[0], fields[1:]...)
	err = cmd.Run() // waits for it to complete
	if err != nil {
		log.Fatal(err)
	}
}

// Execute an external command but don't see its output.
// The command must be given in pieces in a string list.
// That is, split the command manually.
func ExecCmdWithoutOutputOnWindows(fields []string) {
	cmd := exec.Command(fields[0], fields[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Execute an external command in the background.
// Specify just the command, e.g. "gedit" .
// We don't wait for the process to finish.
func ExecCmdInBackground(command string) {
	if IsWindows() {
		// it'll start the command in the background
		fields := []string{"cmd.exe", "/C", "start", "/b", "", command}
		ExecCmdWithoutOutputOnWindows(fields)
	} else {
		newCommand := fmt.Sprintf("sh -c \"%s &\"", command)
		ExecCmdWithoutOutput(newCommand)
	}
}
