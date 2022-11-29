package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jabbalaci/jabbagolib/jos"
	"github.com/jabbalaci/jabbagolib/jpath"
)

func main() {
	var fname string = ""

	if jos.IsLinux() {
		fmt.Println("# we're under Linux")
		fname = "/tmp/test85v6b7c342.txt"
	}

	if jos.IsWindows() {
		fmt.Println("# we're under Windows")
		fname = "C:/tmp/test85v6b7c342.txt"
	}

	fmt.Println("Test file: " + fname)

	os.Remove(fname) // tabula rasa

	// Create file
	new, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	new.Close()

	stats, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	mode := stats.Mode()
	fmt.Printf("Permission File Before: %s\n", mode)

	jpath.MakeExecutable(fname)

	stats, err = os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After:  %s\n", stats.Mode())

	os.Remove(fname)
}
