package jslice

import (
	"reflect"
	"testing"

	"github.com/jabbalaci/jabbagolib/jtest"
)

var AssertBool = jtest.AssertBool
var AssertStr = jtest.AssertStr
var AssertInt = jtest.AssertInt

/////////////////////////////////////////////////////////////////////////////

func TestFindInIntSlice(t *testing.T) {
	li := []int{1, 9, 3, 8, 4, 7, 23, 67}
	AssertInt(FindInIntSlice(li, 1), 0, t)
	AssertInt(FindInIntSlice(li, 8), 3, t)
	AssertInt(FindInIntSlice(li, 23), 6, t)

	AssertInt(FindInIntSlice(li, 5), -1, t)
	AssertInt(FindInIntSlice(li, 2014), -1, t)
}

func TestReverseSliceInt(t *testing.T) {
	AssertBool(reflect.DeepEqual(ReverseIntSlice([]int{}), []int{}), true, t)
	AssertBool(reflect.DeepEqual(ReverseIntSlice([]int{1}), []int{1}), true, t)
	AssertBool(reflect.DeepEqual(ReverseIntSlice([]int{1, 2}), []int{2, 1}), true, t)
	AssertBool(reflect.DeepEqual(ReverseIntSlice([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}), true, t)
}
