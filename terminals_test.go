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

func TestIntSum(t *testing.T) {
	slice := []int{1, 2, 3}
	res := Sum(slice)
	expect := 6

	if res != expect {
		t.Errorf("TestIntSum was incorrect, got: %d, expected: %d", res, expect)
	}
}

func TestUintptrSum(t *testing.T) {
	slice := []uintptr{1, 2, 3}
	res := Sum(slice)
	expect := uintptr(6)

	if res != expect {
		t.Errorf("TestUintptrSum was incorrect, got: %d, expected: %d", res, expect)
	}
}

func TestFloatSum(t *testing.T) {
	slice := []float64{0.668, 0.666, 0.666}
	res := Sum(slice)
	expect := 2.0

	if res != expect {
		t.Errorf("TestFloatSum was incorrect, got: %f, expected: %f", res, expect)
	}
}

func TestStringSum(t *testing.T) {
	slice := []string{"a", "b", "c"}
	res := Sum(slice)
	expect := "abc"

	if res != expect {
		t.Errorf("TestStringSum was incorrect, got: %s, expected: %s", res, expect)
	}
}

func TestByteSum(t *testing.T) {
	slice := []byte{1, 2, 3}
	res := Sum(slice)
	expect := byte(6)

	if res != expect {
		t.Errorf("TestByteSum was incorrect, got: %d, expected: %d", res, expect)
	}
}
