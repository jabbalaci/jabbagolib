package jslice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestFindInIntSlice(t *testing.T) {
	li := []int{1, 9, 3, 8, 4, 7, 23, 67}
	assert.Equal(t, FindInIntSlice(li, 1), 0)
	assert.Equal(t, FindInIntSlice(li, 8), 3)
	assert.Equal(t, FindInIntSlice(li, 23), 6)

	assert.Equal(t, FindInIntSlice(li, 5), -1)
	assert.Equal(t, FindInIntSlice(li, 2014), -1)
}

func TestReverseSliceInt(t *testing.T) {
	assert.Equal(t, reflect.DeepEqual(ReverseIntSlice([]int{}), []int{}), true)
	assert.Equal(t, reflect.DeepEqual(ReverseIntSlice([]int{1}), []int{1}), true)
	assert.Equal(t, reflect.DeepEqual(ReverseIntSlice([]int{1, 2}), []int{2, 1}), true)
	assert.Equal(t, reflect.DeepEqual(ReverseIntSlice([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}), true)
}
