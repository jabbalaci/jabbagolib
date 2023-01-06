package jpy

import (
	"github.com/jabbalaci/jabbagolib/jassert"
	"github.com/jabbalaci/jabbagolib/jmath"
)

// Similar to Python's range().
// It returns a list of integers.
func PyRange(nums ...int) []int {
	size := len(nums)
	jassert.Assert(size >= 1 && size <= 3, "PyRange expects 1, 2 or 3 arguments. No more, no less.")
	//
	var (
		start = 0
		end   int
		step  = 1
	)

	switch size {
	case 1:
		end = nums[0]
	case 2:
		start = nums[0]
		end = nums[1]
	case 3:
		start = nums[0]
		end = nums[1]
		step = nums[2]
	}

	jassert.Assert(step != 0, "step must not be zero") // step cannot be 0

	result := make([]int, 0, jmath.Abs((end-start)/step)+1)

	if step > 0 {
		for i := start; i < end; i += step {
			result = append(result, i)
		}
	} else { // if step < 0
		for i := start; i > end; i += step {
			result = append(result, i)
		}
	}

	return result
}
