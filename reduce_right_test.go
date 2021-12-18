package functools

import "testing"

type reduceRightCase[T any, R comparable] struct {
	name    string
	slice   []T
	initial R
	reducer func(R, T) R
	want    R
}

func TestReduceRight(t *testing.T) {
	t.Run("integers", func(t *testing.T) {

		cases := []reduceRightCase[int, int]{
			{
				name:    "addition",
				slice:   []int{1, 2, 3},
				initial: 0,
				reducer: func(a, b int) int { return a + b },
				want:    6,
			},
			{
				name:    "subtraction",
				slice:   []int{1, 2, 3},
				initial: 0,
				reducer: func(a, b int) int { return a - b },
				want:    -6,
			},
			{
				name:    "multiplication",
				slice:   []int{1, 2, 3},
				initial: 1,
				reducer: func(a, b int) int { return a * b },
				want:    6,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := ReduceRight(c.slice, c.initial, c.reducer)

				if got != c.want {
					t.Errorf("got %v, want %v", got, c.want)
				}
			})
		}
	})

	t.Run("integers and floats", func(t *testing.T) {
		cases := []reduceRightCase[int, float64]{
			{
				name:    "division",
				slice:   []int{1, 2, 3},
				initial: 1.0,
				reducer: func(accum float64, curr int) float64 { return float64(curr) / accum },
				want:    1.5,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := ReduceRight(c.slice, c.initial, c.reducer)

				if got != c.want {
					t.Errorf("got %v, want %v", got, c.want)
				}
			})
		}
	})
}
