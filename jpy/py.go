// Package jpy includes some built-in functions of Python.
package jpy

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jassert"
	ifs "github.com/jabbalaci/jabbagolib/jinterfaces"
)

// Returns the sum of the numbers.
func Sum[T ifs.NumberTypes](numbers []T) T {
	var total T = 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func Min[T ifs.Ordered](li []T) T {
	jassert.Assert(len(li) > 0, "cannot find the minimum element in an empty list")
	//
	mini := li[0]
	for _, curr := range li[1:] {
		if curr < mini {
			mini = curr
		}
	}
	return mini
}

func Max[T ifs.Ordered](li []T) T {
	jassert.Assert(len(li) > 0, "cannot find the maximum element in an empty list")
	//
	maxi := li[0]
	for _, curr := range li[1:] {
		if curr > maxi {
			maxi = curr
		}
	}
	return maxi
}

// Convert a decimal number to binary.
// The prefix '0b' is present.
func Bin(n int) string {
	return fmt.Sprintf("0b%b", n)
}

// Convert a decimal number to octal.
// The prefix '0o' is present.
func Oct(n int) string {
	return fmt.Sprintf("0o%o", n)
}

// Convert a decimal number to hexa.
// The prefix '0x' is present.
func Hex(n int) string {
	return fmt.Sprintf("0x%x", n)
}
