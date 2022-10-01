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
