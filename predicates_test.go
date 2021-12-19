package functools

import "testing"

// All tests
func TestGeqAll(t *testing.T) {
	slice := []int{100, 25, 20, 31, 30}
	if All(slice, func(val int) bool { return val >= 21 }) {
		t.Errorf("TestGeqAll with %v was incorrect, got: %v, expected: %v", slice, true, false)
	}
	slice[2] = 21
	if !All(slice, func(val int) bool { return val >= 21 }) {
		t.Errorf("TestGeqAll with %v was incorrect, got: %v, expected: %v", slice, false, true)
	}
}

// Any tests
func TestLtAny(t *testing.T) {
	slice := []int{20, 31, 22}

	if !Any(slice, func(val int) bool { return val < 21 }) {
		t.Errorf("TestLtAny with %v was incorrect, got: %v, expected: %v", slice, false, true)
	}
	slice[0] = 21
	if Any(slice, func(val int) bool { return val < 21 }) {
		t.Errorf("TestLtAny with %v was incorrect, got: %v, expected: %v", slice, true, false)
	}
}
