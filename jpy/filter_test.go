package jpy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestFilter(t *testing.T) {
	words := []string{"war", "cup", "water", "tree", "storm"}

	expected := []string{"war", "water"}
	got := Filter(words, func(s string) bool {
		return strings.HasPrefix(s, "w")
	})
	assert.Equal(t, got, expected)
	//
	got = Filter(words, func(s string) bool {
		return strings.HasPrefix(s, "xxx")
	})
	assert.Equal(t, got, []string{})
}
