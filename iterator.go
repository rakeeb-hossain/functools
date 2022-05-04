package functools

import (
	"golang.org/x/exp/constraints"
)

type Spliterator[T any] struct {
	tryAdvance       func(func(T)) bool
	forEachRemaining func(func(T))
	trySplit         func() (Spliterator[T], bool)
	characteristics  uint
}

func EmptyIter[T any]() (res Spliterator[T]) {
	res.tryAdvance = func(func(T)) bool {
		return false
	}
	res.forEachRemaining = func(func(T)) {}
	res.trySplit = func() (r Spliterator[T], b bool) {
		return r, b
	}
	return res
}

func sliceIterRec[T any, A ~[]T](slice A, lo int, hi int) (res Spliterator[T]) {
	res.tryAdvance = func(fn func(T)) bool {
		if lo >= hi {
			return false
		}
		fn(slice[lo])
		lo++
		return true
	}

	res.forEachRemaining = func(fn func(T)) {
		for ; lo < hi; lo++ {
			fn(slice[lo])
		}
	}

	res.trySplit = func() (s Spliterator[T], b bool) {
		mid := (hi-lo)/2 + lo
		if mid != lo {
			s, b = sliceIterRec[T, A](slice, mid, hi), true
			// Modify current sliceIter before returning
			hi = mid
		}
		return s, b
	}
	return res
}

func SliceIter[T any, A ~[]T](slice A) (res Spliterator[T]) {
	sliceIterRec[T, A](slice, 0, len(slice))
	return res
}

func RuleIter[T any, A ~func() (T, bool)](rule A) (res Spliterator[T]) {
	res.tryAdvance = func(fn func(T)) bool {
		r, b := rule()
		if b {
			fn(r)
		}
		return b
	}
	res.forEachRemaining = func(fn func(T)) {
		for r, b := rule(); b; r, b = rule() {
			fn(r)
		}
	}
	res.trySplit = func() (s Spliterator[T], b bool) {
		return s, b
	}
	return res
}

func sign[T constraints.Integer](x T) int8 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func RangeIter[T constraints.Integer](start, stop T, step T) (res Spliterator[T]) {
	// Check to ensure no infinite-loop
	if sign(stop-start)*sign(step) < 0 {
		return EmptyIter[T]()
	}

	res.tryAdvance = func(fn func(T)) bool {
		if start >= stop {
			return false
		}
		fn(start)
		start += step
		return true
	}
	res.forEachRemaining = func(fn func(T)) {
		for ; start < stop; start += step {
			fn(start)
		}
	}
	res.trySplit = func() (s Spliterator[T], b bool) {
		mid := (stop-start)/2 + start
		if mid != start {
			s, b = RangeIter[T](mid, stop, step), true
			// Modify stop for this iter
			stop = mid
		}
		return s, b
	}
	return res
}

func ChanIter[T any, C ~chan T](ch C) (res Spliterator[T]) {
	res.tryAdvance = func(fn func(T)) bool {
		v, ok := <-ch
		if ok {
			fn(v)
		}
		return ok
	}
	res.forEachRemaining = func(fn func(T)) {
		for elem := range ch {
			fn(elem)
		}
	}
	res.trySplit = func() (s Spliterator[T], b bool) {
		return s, b
	}
	return res
}

// Iterator is a generic iterator on a slice that lazily evaluates the next element in the slice.
// This is used to lazily evaluate a slice's next value, allowing several applications of functional
// methods on a single list while only incurring a O(1) memory overhead.
type Iterator[T any] func() (T, bool)

// Iter consumes a generic slice and generates a forward-advancing Iterator
//
// Iter is passed a copy of the slice. This does not copy the contents of the slice, but the size of
// the slice is fixed. Therefore, modifications of element the internal slice will affect the Iterator
func Iter[T any, A ~[]T](slice A) Iterator[T] {
	index := 0
	return func() (t T, b bool) {
		if index >= len(slice) {
			return t, b // b is false here
		}
		index++
		return slice[index-1], true
	}
}

// ReverseIter consumes a generic slice and generates a reverse-advancing Iterator
func ReverseIter[T any, A ~[]T](slice A) Iterator[T] {
	index := len(slice) - 1
	return func() (t T, b bool) {
		if index < 0 || index >= len(slice) {
			return t, b // b is false here
		}
		index--
		return slice[index+1], true
	}
}

// Slice converts a generic Iterator to a slice of the appropriate type
func Slice[T any](iter Iterator[T]) []T {
	res := make([]T, 0)
	for val, ok := iter(); ok; val, ok = iter() {
		res = append(res, val)
	}
	return res
}

// Next is an alias for advancing the Iterator
func Next[T any](iter Iterator[T]) (T, bool) {
	return iter()
}
