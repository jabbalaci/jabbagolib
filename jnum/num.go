// Package jnum faciliates work with numbers.
package jnum

import (
	"bytes"
	"fmt"
	"github.com/jabbalaci/jabbagolib/jmisc"
	"github.com/jabbalaci/jabbagolib/jtext"
	"math/big"
	"strconv"
)

var p = fmt.Println
var pf = fmt.Printf
var spf = fmt.Sprintf

// Pow raises a to the power of b.
// The result is a big int.
func Pow(a, b int64) *big.Int {
	jmisc.Assert(!(a == 0 && b == 0), "0 to the power of 0 is undefined")
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

// IsPrime decides whether a number is prime or not.
// Thanks to ewwFatties @reddit for the feedback.
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

// NextPrime returns the next prime
// that is >= than the given parameter.
//
// Usage:
//    np := NextPrime(5)
//    for i := 0; i < 10; i++ {
//    	p(np())
//    }
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

// PrettyNum prettifies a number,
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
