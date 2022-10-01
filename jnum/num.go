// Package jnum faciliates work with numbers.
package jnum

import (
	"bytes"
	"math/big"
	"strconv"

	"github.com/jabbalaci/jabbagolib/jassert"
	"github.com/jabbalaci/jabbagolib/jtext"
)

// Raises `a` to the power of `b`.
// The result is a big int.
func Pow(a, b int64) *big.Int {
	jassert.Assert(!(a == 0 && b == 0), "0 to the power of 0 is undefined")
	if b == 0 {
		return big.NewInt(1)
	}
	//
	n := big.NewInt(a)
	i := int64(1)
	mul := big.NewInt(a)
	for i < b {
		n = n.Mul(n, mul)
		i++
	}
	return n
}

// Decides whether a number is prime or not.
// This is not an efficient implementation.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Returns the next prime
// that is >= the given parameter.
//
// Usage:
//
//	np := NextPrime(5)
//	for i := 0; i < 10; i++ {
//		fmt.Println(np())
//	}
func NextPrime(n int) func() int {
	curr := n
	if curr < 1 {
		curr = 2
	} else if (curr > 2) && (curr%2 == 0) {
		curr++
	}
	return func() int {
		for {
			if IsPrime(curr) {
				result := curr
				if curr > 2 {
					curr += 2
				} else {
					curr++
				}
				return result
			}
			if curr > 2 {
				curr += 2
			} else {
				curr++
			}
		}
	}
}

// Prettifies a number,
// e.g. 1977 => "1,977", or 1234567 => "1,234,567".
func PrettyNum(n int) string {
	neg := n < 0
	if neg {
		n *= -1 // absolute value
	}
	var buffer bytes.Buffer
	s := strconv.Itoa(n)
	for i, v := range jtext.ReverseStr(s) {
		if (i > 0) && (i%3 == 0) {
			buffer.WriteRune(',')
		}
		buffer.WriteRune(v)
	}
	if neg {
		buffer.WriteRune('-')
	}
	return jtext.ReverseStr(buffer.String())
}
