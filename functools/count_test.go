package functools

import "testing"

func TestGeqCount(t *testing.T) {
	slice := []int{1, 100, 200, 3, 14, 21, 32}
	res := Count(slice, func(val int) bool { return val >= 21 })
	expected := 4

	if res != expected {
		t.Errorf("TestLtAny with %v was incorrect, got: %d, expected: %d", slice, res, expected)
	}
}