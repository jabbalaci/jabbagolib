package jrand

import (
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

func TestShuffleSliceInt(t *testing.T) {
	var li []int
	//
	li = []int{}
	Shuffle(li)
	assert.Equal(t, len(li), 0)
	//
	li = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Shuffle(li)
	assert.Equal(t, len(li), 9)
}
