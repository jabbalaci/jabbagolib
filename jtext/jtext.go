package jtext

import (
	"bytes"
	"strconv"
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
)

// Is the given string a palindrome?
func IsPalindrome(s string) bool {
	return s == ReverseStr(s)
}

// Reverse a (Unicode) string.
// from http://stackoverflow.com/questions/1752414
func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// join elements of a []int to a string
func SliceIntToStr(li []int) string {
	var buffer bytes.Buffer
	for i := range li {
		buffer.WriteString(strconv.Itoa(li[i]))
	}
	return buffer.String()
}

// swap the case of each character, e.g. "gOLANG" => "Golang"
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
