package jnum

import (
	"github.com/jabbalaci/jabbagolib/jslice"
	"github.com/jabbalaci/jabbagolib/jtest"
	"reflect"
	"testing"
)

var AssertBool = jtest.AssertBool
var AssertStr = jtest.AssertStr
var AssertInt = jtest.AssertInt

/////////////////////////////////////////////////////////////////////////////

func TestPow(t *testing.T) {
	AssertStr(Pow(2, 0).String(), "1", t)
	AssertStr(Pow(2, 1).String(), "2", t)
	AssertStr(Pow(2, 2).String(), "4", t)
	AssertStr(Pow(2, 10).String(), "1024", t)
	AssertStr(Pow(2, 128).String(), "340282366920938463463374607431768211456", t)
	AssertStr(Pow(9, 9).String(), "387420489", t)
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for i := 0; i <= 20; i++ {
		res := IsPrime(i)
		AssertBool(res, jslice.FindInIntSlice(primes, i) > -1, t)
	}
}

func TestNextPrime(t *testing.T) {
	res := []int{5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	np := NextPrime(5)
	li := []int{}
	for i := 0; i < 10; i++ {
		li = append(li, np())
	}
	AssertBool(reflect.DeepEqual(res, li), true, t)
}

func TestPrettyNum(t *testing.T) {
	AssertStr(PrettyNum(0), "0", t)
	AssertStr(PrettyNum(1), "1", t)
	AssertStr(PrettyNum(100), "100", t)
	AssertStr(PrettyNum(1000), "1,000", t)
	AssertStr(PrettyNum(1977), "1,977", t)
	AssertStr(PrettyNum(1977654), "1,977,654", t)
	//
	AssertStr(PrettyNum(-1), "-1", t)
	AssertStr(PrettyNum(-100), "-100", t)
	AssertStr(PrettyNum(-1000), "-1,000", t)
	AssertStr(PrettyNum(-1977), "-1,977", t)
	AssertStr(PrettyNum(-1977654), "-1,977,654", t)
}
