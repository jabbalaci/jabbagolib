package jpy

// Map operation. Transform each element.
func Map[T, U any](data []T, f func(T) U) []U {
	result := make([]U, len(data))
	for i, e := range data {
		result[i] = f(e)
	}
	return result
}
