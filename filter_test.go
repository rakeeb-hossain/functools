package functools

import "testing"

func TestGeqFilter(t *testing.T) {
	slice := []user{{17}, {21}, {18}, {32}, {49}, {76}}
	geqTwentyOne := func(u user) bool { return u.age >= 21 }
	res := Map(Filter(slice, geqTwentyOne), func(u user) int { return u.age })
	expect := []int{21, 32, 49, 76}

	for i, _ := range res {
		if i >= len(expect) || res[i] != expect[i] {
			t.Errorf("TestGeqFilter was incorrect, got: %v, expected: %v", res, expect)
			return
		}
	}
}
