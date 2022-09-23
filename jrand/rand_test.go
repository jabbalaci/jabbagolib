package jrand

import (
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestSeed(t *testing.T) {
	const LIMIT = 10
	set := mapset.NewSet[int64]()
	for i := 0; i < LIMIT; i++ {
		set.Add(Seed())
	}
	assert.Equal(t, set.Cardinality(), LIMIT)
}

func TestRandInt(t *testing.T) {
	const (
		LO  = 0
		HI  = 10
		REP = 1000
	)
	for i := 0; i < REP; i++ {
		val := RandInt(LO, HI)
		assert.Equal(t, val >= LO && val <= HI, true)
	}
}

func TestRandRange(t *testing.T) {
	const (
		LO  = 0
		HI  = 10
		REP = 1000
	)
	for i := 0; i < REP; i++ {
		val := RandRange(LO, HI)
		assert.Equal(t, val >= LO && val < HI, true)
	}
}

func TestShuffle(t *testing.T) {
	a := []int{}
	Shuffle(a)
	assert.Equal(t, len(a), 0)
	//
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy := append([]int{}, original...)
	cnt := 0
	for i := 0; i < 10; i++ {
		Shuffle(copy)
		if !reflect.DeepEqual(original, copy) {
			cnt++
		}
	}
	assert.Equal(t, cnt > 0, true)
}
