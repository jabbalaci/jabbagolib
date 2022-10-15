package jpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestSumWithInt(t *testing.T) {
	var numbers []int

	numbers = []int{1, 2, 3, 4}
	assert.Equal(t, Sum(numbers), 10)
	//
	numbers = []int{1}
	assert.Equal(t, Sum(numbers), 1)
	//
	numbers = []int{}
	assert.Equal(t, Sum(numbers), 0)
}

func TestSumWithFloat64(t *testing.T) {
	var numbers []float64

	numbers = []float64{1.1, 2.2, 3.3, 4.5}
	assert.Equal(t, Sum(numbers), 11.1)
	//
	numbers = []float64{1.9}
	assert.Equal(t, Sum(numbers), 1.9)
	//
	numbers = []float64{}
	assert.Equal(t, Sum(numbers), 0.0)
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

func TestAbs(t *testing.T) {
	assert.Equal(t, Abs(-5), 5)
	assert.Equal(t, Abs(0), 0)
	assert.Equal(t, Abs(12), 12)
}

func TestBin(t *testing.T) {
	assert.Equal(t, Bin(0), "0b0")
	assert.Equal(t, Bin(123), "0b1111011")
	assert.Equal(t, Bin(576), "0b1001000000")

	assert.Equal(t, Bin(-9), "0b-1001")
}

func TestOct(t *testing.T) {
	assert.Equal(t, Oct(0), "0o0")
	assert.Equal(t, Oct(27), "0o33")
}

func TestHex(t *testing.T) {
	assert.Equal(t, Hex(0), "0x0")
	assert.Equal(t, Hex(27), "0x1b")
}
