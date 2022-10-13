package jslice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestIndex(t *testing.T) {
	li := []int{1, 9, 3, 8, 4, 7, 23, 67}
	assert.Equal(t, Index(li, 1), 0)
	assert.Equal(t, Index(li, 8), 3)
	assert.Equal(t, Index(li, 23), 6)

	assert.Equal(t, Index(li, 5), -1)
	assert.Equal(t, Index(li, 2014), -1)
}

func TestContains(t *testing.T) {
	li := []int{1, 9, 3, 8, 4, 7, 23, 67}
	assert.Equal(t, Contains(li, 1), true)
	assert.Equal(t, Contains(li, 8), true)
	assert.Equal(t, Contains(li, 23), true)

	assert.Equal(t, Contains(li, 5), false)
	assert.Equal(t, Contains(li, 2014), false)
}

func TestReverse(t *testing.T) {
	assert.Equal(t, reflect.DeepEqual(Reverse([]int{}), []int{}), true)
	assert.Equal(t, reflect.DeepEqual(Reverse([]int{1}), []int{1}), true)
	assert.Equal(t, reflect.DeepEqual(Reverse([]int{1, 2}), []int{2, 1}), true)
	assert.Equal(t, reflect.DeepEqual(Reverse([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}), true)
}

func TestSortedInts(t *testing.T) {
	var original, result []int

	original = []int{4, 3, 2, 1}
	result = SortedInts(original)
	assert.Equal(t, reflect.DeepEqual(original, []int{4, 3, 2, 1}), true)
	assert.Equal(t, reflect.DeepEqual(result, []int{1, 2, 3, 4}), true)
	//
	original = []int{}
	result = SortedInts(original)
	assert.Equal(t, reflect.DeepEqual(original, []int{}), true)
	assert.Equal(t, reflect.DeepEqual(result, []int{}), true)
}

func TestSortedFloat64s(t *testing.T) {
	var original, result []float64

	original = []float64{4.1, 3.1, 2.1, 1.1}
	result = SortedFloat64s(original)
	assert.Equal(t, reflect.DeepEqual(original, []float64{4.1, 3.1, 2.1, 1.1}), true)
	assert.Equal(t, reflect.DeepEqual(result, []float64{1.1, 2.1, 3.1, 4.1}), true)
	//
	original = []float64{}
	result = SortedFloat64s(original)
	assert.Equal(t, reflect.DeepEqual(original, []float64{}), true)
	assert.Equal(t, reflect.DeepEqual(result, []float64{}), true)
}

func TestSortedStrings(t *testing.T) {
	var original, result []string

	original = []string{"dd", "bb", "cc", "aa"}
	result = SortedStrings(original)
	assert.Equal(t, reflect.DeepEqual(original, []string{"dd", "bb", "cc", "aa"}), true)
	assert.Equal(t, reflect.DeepEqual(result, []string{"aa", "bb", "cc", "dd"}), true)
	//
	original = []string{}
	result = SortedStrings(original)
	assert.Equal(t, reflect.DeepEqual(original, []string{}), true)
	assert.Equal(t, reflect.DeepEqual(result, []string{}), true)
}
