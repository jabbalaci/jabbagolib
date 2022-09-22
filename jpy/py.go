// Package jpy includes some built-in functions of Python.
package py

type NumberTypes interface {
	int | float64
}

func sum[T NumberTypes](numbers []T) T {
	var total T = 0
	for _, v := range numbers {
		total += v
	}
	return total
}
