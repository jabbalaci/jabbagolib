package jrand

import (
	"github.com/deckarep/golang-set"
	"github.com/jabbalaci/jabbagolib/jtest"
	"testing"
)

var AssertBool = jtest.AssertBool
var AssertStr = jtest.AssertStr
var AssertInt = jtest.AssertInt

/////////////////////////////////////////////////////////////////////////////

func TestSeed(t *testing.T) {
	const LIMIT = 10
	set := mapset.NewSet()
	for i := 0; i < LIMIT; i++ {
		set.Add(Seed())
	}
	AssertInt(set.Cardinality(), LIMIT, t)
}

func TestRandInt(t *testing.T) {
	const (
		LO  = 0
		HI  = 10
		REP = 100
	)
	for i := 0; i < REP; i++ {
		val := RandInt(LO, HI)
		AssertBool(val >= LO && val <= HI, true, t)
	}
}

func TestShuffleSliceInt(t *testing.T) {
	var li []int
	//
	li = []int{}
	ShuffleSliceInt(li)
	AssertInt(len(li), 0, t)
	//
	li = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ShuffleSliceInt(li)
	AssertInt(len(li), 9, t)
}
