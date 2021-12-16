package functools

import "testing"

func TestLtAny(t *testing.T) {
	slice := []int{20, 31, 22}

	if !Any(slice, func(val int) bool {return val < 21}) {
		t.Errorf("TestLtAny with %v was incorrect, got: %v, expected: %v", slice, false, true)
	}
	slice[0] = 21
	if Any(slice, func(val int) bool {return val < 21}) {
		t.Errorf("TestLtAny with %v was incorrect, got: %v, expected: %v", slice, true, false)
	}
}