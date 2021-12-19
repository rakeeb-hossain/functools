package functools

import "testing"

func TestAddMap(t *testing.T) {
	slice := []int{1, 2, 3}
	adder := func(val int) int { return val + 1 }
	res := Map(slice, adder)
	expect := []int{2, 3, 4}

	for i, _ := range res {
		if i >= len(expect) || res[i] != expect[i] {
			t.Errorf("TestAddMap was incorrect, got: %v, expected: %v", res, expect)
			return
		}
	}
}

func TestUserMap(t *testing.T) {
	slice := []user{{32}, {29}, {42}}
	ageTransformer := func(val user) int { return val.age }
	res := Map[user, []user, int](slice, ageTransformer)
	expect := []int{32, 29, 42}

	for i, _ := range res {
		if i >= len(expect) || res[i] != expect[i] {
			t.Errorf("TestUserMap was incorrect, got: %v, expected: %v", res, expect)
			return
		}
	}
}
