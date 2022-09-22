package jtext

import (
	"testing"

	"github.com/jabbalaci/jabbagolib/jtest"
)

var AssertBool = jtest.AssertBool
var AssertStr = jtest.AssertStr

/////////////////////////////////////////////////////////////////////////////

func TestIsPalindrome(t *testing.T) {
	AssertBool(IsPalindrome(""), true, t)
	AssertBool(IsPalindrome("G"), true, t)
	AssertBool(IsPalindrome("Go"), false, t)
	AssertBool(IsPalindrome("görög"), true, t)
	AssertBool(IsPalindrome("radar"), true, t)
	AssertBool(IsPalindrome("Radar"), false, t)
}

func TestReverseStr(t *testing.T) {
	AssertStr(ReverseStr(""), "", t)
	AssertStr(ReverseStr("G"), "G", t)
	AssertStr(ReverseStr("Go"), "oG", t)
	AssertStr(ReverseStr("görög"), "görög", t)
	AssertStr(ReverseStr("László"), "ólzsáL", t)
}

func TestSliceIntToStr(t *testing.T) {
	AssertStr(SliceIntToStr([]int{}), "", t)
	AssertStr(SliceIntToStr([]int{1}), "1", t)
	AssertStr(SliceIntToStr([]int{1, 2, 3}), "123", t)
	AssertStr(SliceIntToStr([]int{-1, -2, -3}), "-1-2-3", t)
	AssertStr(SliceIntToStr([]int{1977, 1980, 1983}), "197719801983", t)
}

func TestSwapCase(t *testing.T) {
	AssertStr(SwapCase(""), "", t)
	AssertStr(SwapCase("g"), "G", t)
	AssertStr(SwapCase("G"), "g", t)
	AssertStr(SwapCase("Go"), "gO", t)
	AssertStr(SwapCase("Golang"), "gOLANG", t)
	AssertStr(SwapCase("László"), "lÁSZLÓ", t)
}
