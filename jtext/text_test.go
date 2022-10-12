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

func TestCapitalize(t *testing.T) {
	assert.Equal(t, Capitalize(""), "")
	assert.Equal(t, Capitalize("g"), "G")
	assert.Equal(t, Capitalize("G"), "G")
	assert.Equal(t, Capitalize("go"), "Go")
	assert.Equal(t, Capitalize("GoLaNG"), "Golang")
	assert.Equal(t, Capitalize("lÁsZlÓ"), "László")
	assert.Equal(t, Capitalize("éVA"), "Éva")
}

func TestCapitalizeAscii(t *testing.T) {
	assert.Equal(t, CapitalizeAscii(""), "")
	assert.Equal(t, CapitalizeAscii("g"), "G")
	assert.Equal(t, CapitalizeAscii("G"), "G")
	assert.Equal(t, CapitalizeAscii("go"), "Go")
	assert.Equal(t, CapitalizeAscii("GoLaNG"), "Golang")
}

func TestCapitalizeAndCapitalizeAscii(t *testing.T) {
	assert.Equal(t, CapitalizeAscii(""), Capitalize(""))
	assert.Equal(t, CapitalizeAscii("g"), Capitalize("G"))
	assert.Equal(t, CapitalizeAscii("G"), Capitalize("G"))
	assert.Equal(t, CapitalizeAscii("go"), Capitalize("Go"))
	assert.Equal(t, CapitalizeAscii("GoLaNG"), Capitalize("Golang"))
}

func TestLevDist(t *testing.T) {
	assert.Equal(t, LevDist("abc", "abc"), 0)
	assert.Equal(t, LevDist("abc", "abcd"), 1)
	assert.Equal(t, LevDist("cat", "Kate"), 2)
	assert.Equal(t, LevDist("Prog. 1", "Prog. 2"), 1)
	assert.Equal(t, LevDist("toned", "roses"), 3)
}

func TestCenter(t *testing.T) {
	assert.Equal(t, Center("*", 4), " *  ")
	assert.Equal(t, Center("*", 3), " * ")
	assert.Equal(t, Center("*", 2), "* ")
	assert.Equal(t, Center("*", 1), "*")
	assert.Equal(t, Center("*", 0), "*")
	assert.Equal(t, Center("*", -2), "*")
	assert.Equal(t, Center("*", 10), "    *     ")
	assert.Equal(t, Center("ABC", 10), "   ABC    ")
}
