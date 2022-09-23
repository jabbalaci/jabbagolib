// Package jslice facilitates the work with slices.
package jslice

// Returns the index of an element.
// If not found, -1 is returned.
func Index[T comparable](haystack []T, needle T) int {
	for i, curr := range haystack {
		if curr == needle {
			return i
		}
	}
	return -1
}

// Returns true if the element is found in the slice.
// Returns false if not found.
func Contains[T comparable](haystack []T, needle T) bool {
	return Index(haystack, needle) > -1
}

// Reverses the order of elements.
// It modifies the slice in place.
func Reverse[T any](li []T) []T {
	for i, j := 0, len(li)-1; i < j; i, j = i+1, j-1 {
		li[i], li[j] = li[j], li[i]
	}
	return li
}
