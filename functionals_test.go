package functools

//type user struct {
//	age int
//}
//
//// Filter tests
//func TestGeqFilter(t *testing.T) {
//	slice := []user{{17}, {21}, {18}, {32}, {49}, {76}}
//	geqTwentyOne := func(u user) bool { return u.age >= 21 }
//	res := Map(Filter(slice, geqTwentyOne), func(u user) int { return u.age })
//	expect := []int{21, 32, 49, 76}
//
//	for i, _ := range res {
//		if i >= len(expect) || res[i] != expect[i] {
//			t.Errorf("TestGeqFilter was incorrect, got: %v, expected: %v", res, expect)
//			return
//		}
//	}
//
//	mapper := func(val user) int { return val.age }
//	Map(func(val int) int { return val + 1 }, Map(mapper, SliceIter[user](slice)))
//}
//
//// Map tests
//func TestAddMap(t *testing.T) {
//	slice := []int{1, 2, 3}
//	adder := func(val int) int { return val + 1 }
//	res := Map(slice, adder)
//	expect := []int{2, 3, 4}
//
//	for i, _ := range res {
//		if i >= len(expect) || res[i] != expect[i] {
//			t.Errorf("TestAddMap was incorrect, got: %v, expected: %v", res, expect)
//			return
//		}
//	}
//}
//
//func TestAddMapIter(t *testing.T) {
//	slice := Iter([]int{1, 2, 3})
//	adder := func(val int) int { return val + 1 }
//	res := Slice(MapIter(slice, adder))
//	expect := []int{2, 3, 4}
//
//	for i, _ := range res {
//		if i >= len(expect) || res[i] != expect[i] {
//			t.Errorf("TestAddMapIter was incorrect, got: %v, expected: %v", res, expect)
//			return
//		}
//	}
//}
//
//func TestUserMap(t *testing.T) {
//	slice := []user{{32}, {29}, {42}}
//	ageTransformer := func(val user) int { return val.age }
//	res := Map[user, []user, int](slice, ageTransformer)
//	expect := []int{32, 29, 42}
//
//	for i, _ := range res {
//		if i >= len(expect) || res[i] != expect[i] {
//			t.Errorf("TestUserMap was incorrect, got: %v, expected: %v", res, expect)
//			return
//		}
//	}
//}
//
//// Reduce tests
//func TestReduceSum(t *testing.T) {
//	slice := []int{1, 2, 3}
//	adder := func(a, b int) int { return a + b }
//	res := Reduce(slice, 0, adder)
//	expect := 6
//
//	if res != expect {
//		t.Errorf("TestReduceSum was incorrect, got: %d, expected: %d", res, expect)
//	}
//}
//
//func TestReduceUserAge(t *testing.T) {
//	slice := []user{{32}, {29}, {42}}
//	adder := func(accum int, val user) int { return accum + val.age }
//	res := Reduce[user, []user, int](slice, 0, adder)
//	expect := 103
//
//	if res != expect {
//		t.Errorf("TestReduceUserAge was incorrect, got: %d, expected: %d", res, expect)
//	}
//}
//
//// ReduceRight tests
//type reduceRightCase[T any, R comparable] struct {
//	name    string
//	slice   []T
//	initial R
//	reducer func(R, T) R
//	want    R
//}
//
//func TestReduceRight(t *testing.T) {
//	t.Run("integers", func(t *testing.T) {
//
//		cases := []reduceRightCase[int, int]{
//			{
//				name:    "addition",
//				slice:   []int{1, 2, 3},
//				initial: 0,
//				reducer: func(a, b int) int { return a + b },
//				want:    6,
//			},
//			{
//				name:    "subtraction",
//				slice:   []int{1, 2, 3},
//				initial: 0,
//				reducer: func(a, b int) int { return a - b },
//				want:    -6,
//			},
//			{
//				name:    "multiplication",
//				slice:   []int{1, 2, 3},
//				initial: 1,
//				reducer: func(a, b int) int { return a * b },
//				want:    6,
//			},
//		}
//
//		for _, c := range cases {
//			t.Run(c.name, func(t *testing.T) {
//				got := ReduceRight(c.slice, c.initial, c.reducer)
//
//				if got != c.want {
//					t.Errorf("got %v, want %v", got, c.want)
//				}
//			})
//		}
//	})
//
//	t.Run("integers and floats", func(t *testing.T) {
//		cases := []reduceRightCase[int, float64]{
//			{
//				name:    "division",
//				slice:   []int{1, 2, 3},
//				initial: 1.0,
//				reducer: func(accum float64, curr int) float64 { return float64(curr) / accum },
//				want:    1.5,
//			},
//		}
//
//		for _, c := range cases {
//			t.Run(c.name, func(t *testing.T) {
//				got := ReduceRight(c.slice, c.initial, c.reducer)
//
//				if got != c.want {
//					t.Errorf("got %v, want %v", got, c.want)
//				}
//			})
//		}
//	})
//}
//
//// ForEach tests
//func TestClosureForEach(t *testing.T) {
//	slice := []int{1, 2, 3}
//	res := 0
//	ForEach(slice, func(val int) { res += val })
//	expect := 6
//
//	if res != expect {
//		t.Errorf("TestClosureForEach was incorrect, got: %d, expected: %d", res, expect)
//	}
//}
//
//// Count tests
//func TestGeqCount(t *testing.T) {
//	slice := []int{1, 100, 200, 3, 14, 21, 32}
//	res := Count(slice, func(val int) bool { return val >= 21 })
//	expected := 4
//
//	if res != expected {
//		t.Errorf("TestLtAny with %v was incorrect, got: %d, expected: %d", slice, res, expected)
//	}
//}
