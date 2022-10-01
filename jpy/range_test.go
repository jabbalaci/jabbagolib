package jpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestPyRange(t *testing.T) {
	assert.Equal(t, PyRange(0, 5), []int{0, 1, 2, 3, 4})
	assert.Equal(t, PyRange(3, 7), []int{3, 4, 5, 6})
	assert.Equal(t, PyRange(3, 4), []int{3})
	assert.Equal(t, PyRange(3, 3), []int{})
	//
	assert.Equal(t, PyRange(10), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, PyRange(1), []int{0})
	assert.Equal(t, PyRange(-4), []int{})
	//
	assert.Equal(t, PyRange(4, 20, 2), []int{4, 6, 8, 10, 12, 14, 16, 18})
	assert.Equal(t, PyRange(4, 10, 1), []int{4, 5, 6, 7, 8, 9})
	assert.Equal(t, PyRange(10, 4, 1), []int{})
	//
	assert.Equal(t, PyRange(10, 4, -1), []int{10, 9, 8, 7, 6, 5})
}
