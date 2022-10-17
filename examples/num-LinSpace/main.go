package main

import (
	"fmt"
	"runtime"

	"github.com/jabbalaci/jabbagolib/jnum"
)

func main() {
	result := jnum.LinSpace(0, 440_000_000, runtime.NumCPU())
	for _, v := range result {
		fmt.Printf("%.2f\n", v)
	}
}
