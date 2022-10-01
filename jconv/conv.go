package jconv

import "github.com/jabbalaci/jabbagolib/jassert"

func CharToDigit(c byte) int {
	jassert.Assert(c >= '0' && c <= '9', "Invalid character, not a digit")
	//
	return int(c - '0')
}
