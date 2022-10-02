package jconv

import "github.com/jabbalaci/jabbagolib/jassert"

// Convert a byte (char) to digit.
// Example: '2' (byte) -> 2 (int)
func CharToDigit(c byte) int {
	jassert.Assert(c >= '0' && c <= '9', "Invalid character, not a digit")
	//
	return int(c - '0')
}

// Convert a rune to digit.
// Example: '2' (rune) -> 2 (int)
func RuneToDigit(r rune) int {
	var b byte = byte(r)
	return CharToDigit(b)
}
