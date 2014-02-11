package jnum

import (
	"bytes"
	"github.com/jabbalaci/jabbagolib/jtext"
	"math"
	"math/big"
	"strconv"
)

// a to the power of b
// the result is a big int
func Pow(a, b int64) *big.Int {
	n := big.NewInt(a)
	i := int64(1)
	mul := big.NewInt(a)
	for i < b {
		n = n.Mul(n, mul)
		i++
	}
	return n
}

// reverse the order of elements in a []int
func ReverseSliceInt(li []int) []int {
	for i, j := 0, len(li)-1; i < j; i, j = i+1, j-1 {
		li[i], li[j] = li[j], li[i]
	}
	return li
}

// decide whether a number is prime or not
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

	i := 3
	maxi := math.Sqrt(float64(n)) + 1
	for float64(i) <= maxi {
		if n%i == 0 {
			return false
		}
		i += 2
	}

	return true
}

// return the next prime
// usage:
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

// prettify a number, e.g. 1977 => "1,977", or 1234567 => "1,234,567"
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
