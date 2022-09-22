package jtext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, IsPalindrome(""), true)
	assert.Equal(t, IsPalindrome("G"), true)
	assert.Equal(t, IsPalindrome("Go"), false)
	assert.Equal(t, IsPalindrome("görög"), true)
	assert.Equal(t, IsPalindrome("radar"), true)
	assert.Equal(t, IsPalindrome("Radar"), false)
}

func TestReverseStr(t *testing.T) {
	assert.Equal(t, ReverseStr(""), "")
	assert.Equal(t, ReverseStr("G"), "G")
	assert.Equal(t, ReverseStr("Go"), "oG")
	assert.Equal(t, ReverseStr("görög"), "görög")
	assert.Equal(t, ReverseStr("László"), "ólzsáL")
}

func TestSliceIntToStr(t *testing.T) {
	assert.Equal(t, SliceIntToStr([]int{}), "")
	assert.Equal(t, SliceIntToStr([]int{1}), "1")
	assert.Equal(t, SliceIntToStr([]int{1, 2, 3}), "123")
	assert.Equal(t, SliceIntToStr([]int{-1, -2, -3}), "-1-2-3")
	assert.Equal(t, SliceIntToStr([]int{1977, 1980, 1983}), "197719801983")
}

func TestSwapCase(t *testing.T) {
	assert.Equal(t, SwapCase(""), "")
	assert.Equal(t, SwapCase("g"), "G")
	assert.Equal(t, SwapCase("G"), "g")
	assert.Equal(t, SwapCase("Go"), "gO")
	assert.Equal(t, SwapCase("Golang"), "gOLANG")
	assert.Equal(t, SwapCase("László"), "lÁSZLÓ")
}
