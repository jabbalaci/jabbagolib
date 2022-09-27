// Package jpy includes some built-in functions of Python.
package jpy

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}

type NumberTypes interface {
	Integer | Float
}

// --------------------------------------------------------------------------

func sum[T NumberTypes](numbers []T) T {
	var total T = 0
	for _, v := range numbers {
		total += v
	}
	return total
}
