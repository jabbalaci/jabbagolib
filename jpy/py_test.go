package py

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestSumWithInt(t *testing.T) {
	var numbers []int

	numbers = []int{1, 2, 3, 4}
	assert.Equal(t, sum(numbers), 10)
	//
	numbers = []int{1}
	assert.Equal(t, sum(numbers), 1)
	//
	numbers = []int{}
	assert.Equal(t, sum(numbers), 0)
}
