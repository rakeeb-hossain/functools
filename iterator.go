package functools

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
