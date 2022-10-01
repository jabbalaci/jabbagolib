package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jos"
)

func main() {
	fmt.Println("operating system:", jos.GetOperatingSystem())
	fmt.Println()
	fmt.Println("Is it Linux?", jos.IsLinux())
	fmt.Println("Is it Windows?", jos.IsWindows())
	fmt.Println("Is it Mac?", jos.IsMac())
}
