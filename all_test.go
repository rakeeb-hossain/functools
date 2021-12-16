package functools

import "testing"

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
