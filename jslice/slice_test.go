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
