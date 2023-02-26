package jmath

import (
	"reflect"
	"testing"

	"github.com/jabbalaci/jabbagolib/jslice"
	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestPow(t *testing.T) {
	assert.Equal(t, Pow(2, 0).String(), "1")
	assert.Equal(t, Pow(2, 1).String(), "2")
	assert.Equal(t, Pow(2, 2).String(), "4")
	assert.Equal(t, Pow(2, 10).String(), "1024")
	assert.Equal(t, Pow(2, 128).String(), "340282366920938463463374607431768211456")
	assert.Equal(t, Pow(9, 9).String(), "387420489")
	//
	assert.Equal(t, Pow(0, 5).String(), "0")
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for i := 0; i <= 20; i++ {
		result := IsPrime(i)
		assert.Equal(t, jslice.Contains(primes, i), result)
	}
}

func TestNextPrime(t *testing.T) {
	res := []int{5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	np := NextPrime(5)
	li := []int{}
	for i := 0; i < 10; i++ {
		li = append(li, np())
	}
	assert.Equal(t, reflect.DeepEqual(res, li), true)
}

func TestGetPrimesBelow(t *testing.T) {
	li := GetPrimesBelow(10)
	result := []int{2, 3, 5, 7}
	assert.Equal(t, reflect.DeepEqual(li, result), true)
	//
	li = GetPrimesBelow(11) // 11 not included
	result = []int{2, 3, 5, 7}
	assert.Equal(t, reflect.DeepEqual(li, result), true)
	//
	li = GetPrimesBelow(12) // 11 must be included
	result = []int{2, 3, 5, 7, 11}
	assert.Equal(t, reflect.DeepEqual(li, result), true)
	//
	li = GetPrimesBelow(100)
	result = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	assert.Equal(t, len(result), 25)
	assert.Equal(t, reflect.DeepEqual(li, result), true)
	//
	li = GetPrimesBelow(1000)
	assert.Equal(t, len(li), 168)
}

func TestPrettyNum(t *testing.T) {
	assert.Equal(t, PrettyNum(0), "0")
	assert.Equal(t, PrettyNum(1), "1")
	assert.Equal(t, PrettyNum(100), "100")
	assert.Equal(t, PrettyNum(1000), "1,000")
	assert.Equal(t, PrettyNum(1977), "1,977")
	assert.Equal(t, PrettyNum(1977654), "1,977,654")
	//
	assert.Equal(t, PrettyNum(-1), "-1")
	assert.Equal(t, PrettyNum(-100), "-100")
	assert.Equal(t, PrettyNum(-1000), "-1,000")
	assert.Equal(t, PrettyNum(-1977), "-1,977")
	assert.Equal(t, PrettyNum(-1977654), "-1,977,654")
}

func TestSignum(t *testing.T) {
	assert.Equal(t, Signum(-5), -1)
	assert.Equal(t, Signum(0), 0)
	assert.Equal(t, Signum(42), 1)

	assert.Equal(t, Signum(-5.3), -1)
	assert.Equal(t, Signum(0.0), 0)
	assert.Equal(t, Signum(42.5), 1)
}

func TestAbs(t *testing.T) {
	assert.Equal(t, Abs(-5), 5)
	assert.Equal(t, Abs(0), 0)
	assert.Equal(t, Abs(12), 12)
}

func TestProduct(t *testing.T) {
	var numbers []int

	numbers = []int{}
	assert.Equal(t, Product(numbers), 1)
	//
	numbers = []int{1}
	assert.Equal(t, Product(numbers), 1)
	//
	numbers = []int{1, 2, 3}
	assert.Equal(t, Product(numbers), 6)
	//
	numbers = []int{6, 4, 2}
	assert.Equal(t, Product(numbers), 48)
}
