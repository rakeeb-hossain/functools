package functools

import "testing"

func TestClosureForEach(t *testing.T) {
	slice := []int{1, 2, 3}
	res := 0
	ForEach(slice, func(val int) { res += val })
	expect := 6

	if res != expect {
		t.Errorf("TestClosureForEach was incorrect, got: %d, expected: %d", res, expect)
	}
}