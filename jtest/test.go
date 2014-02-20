// Package jtest facilitates writing unit tests.
package jtest

import (
	"runtime/debug"
	"testing"
)

// AssertBool receives a bool and an expected bool.
// If they are not equal, then the test fails.
func AssertBool(got bool, expected bool, t *testing.T) {
	if got != expected {
		debug.PrintStack()
		t.Error("Expected:", expected, "; got:", got)
	}
}

// AssertStr receives a string and an expected string.
// If they are not equal, then the test fails.
func AssertStr(got string, expected string, t *testing.T) {
	if got != expected {
		debug.PrintStack()
		t.Error("Expected:", expected, "; got:", got)
	}
}

// AssertInt receives an int and an expected int.
// If they are not equal, then the test fails.
func AssertInt(got int, expected int, t *testing.T) {
	if got != expected {
		debug.PrintStack()
		t.Error("Expected:", expected, "; got:", got)
	}
}
