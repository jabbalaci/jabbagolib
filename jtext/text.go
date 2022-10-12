// Package jtext facilitates work with strings.
package jtext

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	ASCII_LETTERS   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ASCII_LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	ASCII_UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS          = "0123456789"
	HEXDIGITS       = "0123456789abcdefABCDEF"
	LETTERS         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	OCTDIGITS       = "01234567"
	PRINTABLE       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	PUNCTUATION     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	WHITESPACE      = " \t\n\r\x0b\x0c"
)

// IsPalindrome tests whether the given string is a palindrome or not.
func IsPalindrome(s string) bool {
	return s == ReverseStr(s)
}

// ReverseStr reverses a (Unicode) string.
func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SliceIntToStr joins elements of a []int to a string.
func SliceIntToStr(li []int) string {
	var buffer bytes.Buffer
	for i := range li {
		buffer.WriteString(strconv.Itoa(li[i]))
	}
	return buffer.String()
}

// SwapCase swaps the case of each character,
// e.g. "gOLANG" => "Golang".
func SwapCase(s string) string {
	var buffer bytes.Buffer
	for _, r := range s {
		if unicode.IsLower(r) {
			buffer.WriteRune(unicode.ToUpper(r))
		} else if unicode.IsUpper(r) {
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// Returns the capitalized version of the given string.
// It works with Unicode text.
// Ex.: "GoLaNG" => "Golang", "éVA" => "Éva"
// If your text is plain ASCII, CapitalizeAscii() will be faster.
// You can also use this for ASCII texts, the result will be the same.
func Capitalize(s string) string {
	var buffer bytes.Buffer
	for idx, r := range s {
		if idx == 0 {
			buffer.WriteRune(unicode.ToUpper(r))
		} else {
			buffer.WriteRune(unicode.ToLower(r))
		}
	}
	return buffer.String()
}

// Returns the capitalized version of the given string.
// It works with ASCII text only!
// Ex.: "GoLaNG" => "Golang"
// If your text is Unicode, use Capitalize() instead.
func CapitalizeAscii(s string) string {
	if len(s) == 0 {
		return ""
	}
	// else
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

// Levenshtein distance between two strings.
// Also known as "edit distance".
func LevDist(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

// Centers a text on a given width. Padding is done with spaces.
func Center(s string, width int) string {
	if width <= len(s) {
		return s
	}
	// else
	return fmt.Sprintf("%*s", -width, fmt.Sprintf("%*s", (width+len(s))/2, s))
}
