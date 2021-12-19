package functools

import "testing"

func TestReduceSum(t *testing.T) {
	slice := []int{1, 2, 3}
	adder := func(a, b int) int { return a + b }
	res := Reduce(slice, 0, adder)
	expect := 6

	if res != expect {
		t.Errorf("TestReduceSum was incorrect, got: %d, expected: %d", res, expect)
	}
}

type user struct {
	age int
}

func TestReduceUserAge(t *testing.T) {
	slice := []user{{32}, {29}, {42}}
	adder := func(accum int, val user) int { return accum + val.age }
	res := Reduce[user, []user, int](slice, 0, adder)
	expect := 103

	if res != expect {
		t.Errorf("TestReduceUserAge was incorrect, got: %d, expected: %d", res, expect)
	}
}
