// Package jmath faciliates work with numbers.
package jmath

import (
	"bytes"
	"math"
	"math/big"
	"strconv"

	"github.com/jabbalaci/jabbagolib/jassert"
	ifs "github.com/jabbalaci/jabbagolib/jinterfaces"
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

// Returns all prime numbers below the given threshold.
// The upper limit `n` is not included.
// It uses the sieve of Eratosthenes algorithm.
func GetPrimesBelow(n int) []int {
	result := []int{}
	lst := make([]bool, n)
	for i := range lst {
		lst[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if lst[i] {
			for j := i * i; j < n; j += i {
				lst[j] = false
			}
		}
	}
	for i := 2; i < n; i++ {
		if lst[i] {
			result = append(result, i)
		}
	}

	return result
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

// Returns num evenly spaced intervals, calculated over the interval [lo, hi].
// The endpoint (hi) is included.
func LinSpace(lo, hi float64, num int) []float64 {
	result := make([]float64, 1, num)
	result[0] = lo
	diff := (hi - lo) / float64(num)
	curr := lo
	for i := 0; i < num-1; i++ {
		curr += diff
		result = append(result, curr)
	}
	result = append(result, hi)

	return result
}

// Returns the sign of the given number.
// If it's positive, return +1.
// If it's negative, return -1.
// Otherwise, return 0.
func Signum[T ifs.NumberTypes](number T) int {
	if number < 0 {
		return -1
	} else if number > 0 {
		return 1
	} else {
		return 0
	}
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

// Returns the product of the numbers.
func Product[T ifs.NumberTypes](numbers []T) T {
	var p T = 1
	for _, value := range numbers {
		p *= value
	}
	return p
}
