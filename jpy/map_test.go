package jpy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestMap(t *testing.T) {
	words := []string{"war", "cup", "water"}

	expected := []string{"WAR", "CUP", "WATER"}
	got := Map(words, func(s string) string {
		return strings.ToUpper(s)
	})
	assert.Equal(t, got, expected)
	//
	expected2 := []int{3, 3, 5}
	got2 := Map(words, func(s string) int {
		return len(s)
	})
	assert.Equal(t, got2, expected2)
}
