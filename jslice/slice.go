// Package jslice facilitates the work with slices.
package jslice

// FindInIntSlice gets an int slice and an elem. It
// return the index position of the elem in the slice
// or -1 if the elem is not in the slice.
func FindInIntSlice(li []int, elem int) int {
	for i, curr := range li {
		if curr == elem {
			return i
		}
	}
	return -1
}

// ReverseIntSlice reverses the order of elements in a []int.
func ReverseIntSlice(li []int) []int {
	for i, j := 0, len(li)-1; i < j; i, j = i+1, j-1 {
		li[i], li[j] = li[j], li[i]
	}
	return li
}
