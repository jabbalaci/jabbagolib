package jconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestCharToDigit(t *testing.T) {
	assert.Equal(t, CharToDigit('0'), 0)
	assert.Equal(t, CharToDigit('1'), 1)
	assert.Equal(t, CharToDigit('2'), 2)
	assert.Equal(t, CharToDigit('3'), 3)
	assert.Equal(t, CharToDigit('4'), 4)
	assert.Equal(t, CharToDigit('5'), 5)
	assert.Equal(t, CharToDigit('6'), 6)
	assert.Equal(t, CharToDigit('7'), 7)
	assert.Equal(t, CharToDigit('8'), 8)
	assert.Equal(t, CharToDigit('9'), 9)
}

func TestRuneToDigit(t *testing.T) {
	assert.Equal(t, CharToDigit('0'), 0)
	assert.Equal(t, CharToDigit('1'), 1)
	assert.Equal(t, CharToDigit('2'), 2)
	assert.Equal(t, CharToDigit('3'), 3)
	assert.Equal(t, CharToDigit('4'), 4)
	assert.Equal(t, CharToDigit('5'), 5)
	assert.Equal(t, CharToDigit('6'), 6)
	assert.Equal(t, CharToDigit('7'), 7)
	assert.Equal(t, CharToDigit('8'), 8)
	assert.Equal(t, CharToDigit('9'), 9)
}

func TestIntToStr(t *testing.T) {
	assert.Equal(t, IntToStr(0), "0")
	assert.Equal(t, IntToStr(1), "1")
	assert.Equal(t, IntToStr(2), "2")
	assert.Equal(t, IntToStr(3), "3")
	assert.Equal(t, IntToStr(4), "4")
	assert.Equal(t, IntToStr(5), "5")
	assert.Equal(t, IntToStr(6), "6")
	assert.Equal(t, IntToStr(7), "7")
	assert.Equal(t, IntToStr(8), "8")
	assert.Equal(t, IntToStr(9), "9")
	//
	assert.Equal(t, IntToStr(2020), "2020")
	assert.Equal(t, IntToStr(-2), "-2")
	assert.Equal(t, IntToStr(12345), "12345")
}

func TestStrToInt(t *testing.T) {
	n, err := StrToInt("0")
	assert.Equal(t, n, 0)
	assert.Nil(t, err)
	//
	n, err = StrToInt("1")
	assert.Equal(t, n, 1)
	assert.Nil(t, err)
	//
	n, err = StrToInt("2")
	assert.Equal(t, n, 2)
	assert.Nil(t, err)
	//
	n, err = StrToInt("2020")
	assert.Equal(t, n, 2020)
	assert.Nil(t, err)
	//
	n, err = StrToInt("-2")
	assert.Equal(t, n, -2)
	assert.Nil(t, err)
	//
	n, err = StrToInt("12345")
	assert.Equal(t, n, 12345)
	assert.Nil(t, err)
	//
	_, err = StrToInt("hello")
	assert.NotNil(t, err)
}

func TestFloat64ToStr(t *testing.T) {
	assert.Equal(t, Float64ToStr(3.14), "3.14")
	assert.Equal(t, Float64ToStr(3.145), "3.145")
	assert.Equal(t, Float64ToStr(3.1459), "3.1459")
}

func TestStrToFloat64(t *testing.T) {
	f, err := StrToFloat64("3.14")
	assert.Equal(t, f, 3.14)
	assert.Nil(t, err)
	//
	f, err = StrToFloat64("3.145")
	assert.Equal(t, f, 3.145)
	assert.Nil(t, err)
	//
	f, err = StrToFloat64("3.1459")
	assert.Equal(t, f, 3.1459)
	assert.Nil(t, err)
	//
	_, err = StrToFloat64("hello")
	assert.NotNil(t, err)
}
