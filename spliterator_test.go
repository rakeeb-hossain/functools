package functools

import (
	"reflect"
	"testing"
)

type iterTestCase[T any] struct {
	name  string
	slice []T
}

func TestIter(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		cases := []iterTestCase[int]{
			{
				"integer1",
				[]int{1, 2, 3},
			},
			{
				"integer2",
				[]int{100, 0, -1, -1000},
			},
			{
				"empty slice",
				[]int{},
			},
		}

		for _, c := range cases {
			iter := Iter(c.slice)
			for i, _ := range c.slice {
				val, ok := Next(iter)
				if !ok {
					t.Errorf("got %v, want %v", nil, c.slice[i])
					continue
				}
				if val != c.slice[i] {
					t.Errorf("got %v, want %v", val, c.slice[i])
				}
			}

			// Make sure iter is empty
			val, ok := Next(iter)
			if ok {
				t.Errorf("got %v, want %v", val, nil)
			}
		}
	})

	t.Run("slice of string slices", func(t *testing.T) {
		cases := []iterTestCase[[]string]{
			{
				"empty string slice of slices",
				[][]string{},
			},
			{
				"string1",
				[][]string{{"asdf", "rakeeb", "gopher"}, {"kevin", "trevor"}},
			},
		}

		for _, c := range cases {
			iter := Iter(c.slice)

			for i, _ := range c.slice {
				val, ok := Next(iter)
				if !ok {
					t.Errorf("got %v, want %v", nil, c.slice[i])
				}
				if !reflect.DeepEqual(val, c.slice[i]) {
					t.Errorf("got %v, want %v", val, c.slice[i])
				}
			}

			// Make sure iter is empty
			val, ok := Next(iter)
			if ok {
				t.Errorf("got %v, want %v", val, nil)
			}
		}
	})
}

func TestReverseIter(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		cases := []iterTestCase[int]{
			{
				"integer1",
				[]int{1, 2, 3},
			},
			{
				"integer2",
				[]int{100, 0, -1, -1000},
			},
			{
				"empty slice",
				[]int{},
			},
		}

		for _, c := range cases {
			iter := ReverseIter(c.slice)

			for i := len(c.slice) - 1; i >= 0; i-- {
				val, ok := Next(iter)
				if !ok {
					t.Errorf("got %v, want %v", nil, c.slice[i])
				}
				if val != c.slice[i] {
					t.Errorf("got %v, want %v", val, c.slice[i])
				}
			}

			// Make sure iter is empty
			val, ok := Next(iter)
			if ok {
				t.Errorf("got %v, want %v", val, nil)
			}
		}
	})

	t.Run("slice of string slices", func(t *testing.T) {
		cases := []iterTestCase[[]string]{
			{
				"empty string slice of slices",
				[][]string{},
			},
			{
				"string1",
				[][]string{{"asdf", "rakeeb", "gopher"}, {"kevin", "trevor"}},
			},
		}

		for _, c := range cases {
			iter := ReverseIter(c.slice)

			for i := len(c.slice) - 1; i >= 0; i-- {
				val, ok := Next(iter)
				if !ok {
					t.Errorf("got %v, want %v", nil, c.slice[i])
				}
				if !reflect.DeepEqual(val, c.slice[i]) {
					t.Errorf("got %v, want %v", val, c.slice[i])
				}
			}

			// Make sure iter is empty
			val, ok := Next(iter)
			if ok {
				t.Errorf("got %v, want %v", val, nil)
			}
		}
	})
}

func reverse[T any](slice []T) {
	start := 0
	last := len(slice) - 1
	for start < last {
		slice[start], slice[last] = slice[last], slice[start]
		start++
		last--
	}
}

func TestSlice(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		cases := []iterTestCase[int]{
			{
				"integer1",
				[]int{1, 2, 3},
			},
			{
				"integer2",
				[]int{100, 0, -1, -1000},
			},
			{
				"empty slice",
				[]int{},
			},
		}

		for _, c := range cases {
			slice := Slice(Iter(c.slice))

			if !reflect.DeepEqual(slice, c.slice) {
				t.Errorf("got %v, want %v", slice, c.slice)
			}
		}

		for _, c := range cases {
			slice := Slice(ReverseIter(c.slice))
			reverse(slice)

			if !reflect.DeepEqual(slice, c.slice) {
				t.Errorf("got %v, want %v", slice, c.slice)
			}
		}
	})

	t.Run("slice of string slices", func(t *testing.T) {
		cases := []iterTestCase[[]string]{
			{
				"empty string slice of slices",
				[][]string{},
			},
			{
				"string1",
				[][]string{{"asdf", "rakeeb", "gopher"}, {"kevin", "trevor"}},
			},
		}

		for _, c := range cases {
			slice := Slice(Iter(c.slice))

			if !reflect.DeepEqual(slice, c.slice) {
				t.Errorf("got %v, want %v", slice, c.slice)
			}
		}

		for _, c := range cases {
			slice := Slice(ReverseIter(c.slice))
			reverse(slice)

			if !reflect.DeepEqual(slice, c.slice) {
				t.Errorf("got %v, want %v", slice, c.slice)
			}
		}
	})
}
