package jconv

import (
	"strconv"

	"github.com/jabbalaci/jabbagolib/jassert"
)

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

// Convert an int to string.
// Example: 42 (int) -> "42" (string)
func IntToStr(n int) string {
	return strconv.Itoa(n)
}

// Convert a string to int.
// Example: "42" (string) -> 42 (int)
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// Convert a float64 to string.
// Example: 3.14 -> "3.14"
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Convert a string to float64.
// Example: "3.14159" -> 3.14159
func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
