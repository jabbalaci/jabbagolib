package jpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestSumWithInt(t *testing.T) {
	var numbers []int

	numbers = []int{1, 2, 3, 4}
	assert.Equal(t, sum(numbers), 10)
	//
	numbers = []int{1}
	assert.Equal(t, sum(numbers), 1)
	//
	numbers = []int{}
	assert.Equal(t, sum(numbers), 0)
}

func TestSumWithFloat64(t *testing.T) {
	var numbers []float64

	numbers = []float64{1.1, 2.2, 3.3, 4.5}
	assert.Equal(t, sum(numbers), 11.1)
	//
	numbers = []float64{1.9}
	assert.Equal(t, sum(numbers), 1.9)
	//
	numbers = []float64{}
	assert.Equal(t, sum(numbers), 0.0)
}
