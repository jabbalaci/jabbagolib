// Package jpy includes some built-in functions of Python.
package jpy

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jassert"
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

// Supports the order operators such as > and < .
type Ordered interface {
	Integer | Float | ~string
}

type NumberTypes interface {
	Integer | Float
}

// --------------------------------------------------------------------------

// Returns the sum of the numbers.
func Sum[T NumberTypes](numbers []T) T {
	var total T = 0
	for _, value := range numbers {
		total += value
	}
	return total
}

// Returns the product of the numbers.
func Product[T NumberTypes](numbers []T) T {
	var p T = 1
	for _, value := range numbers {
		p *= value
	}
	return p
}

// Returns the absolute value of the number.
// The input number is an int.
// math.Abs() expects and returns a float64.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min[T Ordered](li []T) T {
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

func Max[T Ordered](li []T) T {
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
// The '0b' prefix is not present.
func Bin(n int) string {
	return fmt.Sprintf("%b", n)
}
