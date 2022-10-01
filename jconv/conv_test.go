package jconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestCharToDigit(t *testing.T) {
	assert.Equal(t, CharToDigit('0'), 0)
	assert.Equal(t, CharToDigit('1'), 1)
	assert.Equal(t, CharToDigit('2'), 2)
	assert.Equal(t, CharToDigit('3'), 3)
	assert.Equal(t, CharToDigit('4'), 4)
	assert.Equal(t, CharToDigit('5'), 5)
	assert.Equal(t, CharToDigit('6'), 6)
	assert.Equal(t, CharToDigit('7'), 7)
	assert.Equal(t, CharToDigit('8'), 8)
	assert.Equal(t, CharToDigit('9'), 9)
}
