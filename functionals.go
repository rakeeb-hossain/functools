// Contains classic generic functional methods

package functools

// Map consumes a slice of a generic type and returns a slice with the supplied mapping function
// applied to each element.
//
// mapper should be error-safe. It should handle any errors internally and return the desired type.
// If other arguments are required by mapper, mapper should be made a closure with the appropriate
// variables referenced.
func Map[A any, B any](mapper func(A) B, iter Spliterator[A]) (res Spliterator[B]) {
	res.tryAdvance = func(fn func(B)) bool {
		_mapper := func(a A) {
			fn(mapper(a))
		}
		return iter.tryAdvance(_mapper)
	}
	res.forNextK = func(k int, fn func(B)) {
		_mapper := func(a A) {
			fn(mapper(a))
		}
		iter.forNextK(k, _mapper)
	}
	res.trySplit = iter.trySplit
}

// Stateful op

func Sorted[T any](iter Spliterator[T]) (res Spliterator[T]) {
	_buffer := make([]T, 1)
	res.tryAdvance = func(fn func(T)) bool {

	}
}

//func ChunkIter[T any](iter Iterator[T], len int) Iterator[[]T] {
//	return func() (lst []T, b bool) {
//		res, ok := Next(iter)
//		if !ok {
//			return lst, ok
//		}
//
//		lst = make([]T, len)
//		lst[0] = res
//		for i := 1; i < len; i++ {
//			res, ok := Next(iter)
//			if !ok {
//				return lst, true
//			}
//			lst[i] = res
//		}
//		return lst, true
//	}
//}

// Filter consumes a slice of a generic type and returns a slice with only the elements which returned true after
// applying the predicate.
//
// Elements are returned in the same order that they were supplied in the slice.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func Filter[T any, A ~[]T](slice A, predicate func(T) bool) A {
	res := make(A, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func FilterIter[T any](iter Iterator[T], predicate func(T) bool) Iterator[T] {
	return func() (t T, b bool) {
		for val, ok := Next(iter); ok; val, ok = Next(iter) {
			if !ok || predicate(val) {
				return val, ok
			}
		}
		return t, b // b is false here
	}
}

// Reduce consumes a slice of a generic type and an initial value. It reduces the slice to a single value by applying
// the binary reducer function to each successive element in the slice.
//
// Vacuously, empty slices return the initial value provided.
//
// reducer should be error-safe. It should handle any errors internally and return the desired type.
// If other arguments are required by reducer, reducer should be made a closure with the appropriate
// variables referenced.
func Reduce[T any, R any](iter Spliterator[T], initial R, reducer func(R, T) R) R {
	accum := initial
	iter.forNextK(-1, func(v T) {
		accum = reducer(accum, v)
	})
	return accum
}

// ReduceRight consumes a slice of a generic type and an initial value. It
// reduces the slice to a single value by applying the binary reducer function
// to each element in the slice. ReduceRight differs from Reduce by iterating
// from the last element to the first element.
//
// Vacuously, empty slices return the initial value provided.
//
// reducer should error-safe. It should handle any errors internally and return
// the desired type. If other arguments are required by reducer, reducer should
// be made a closure with the appropriate variables referenced.
func ReduceRight[T any, A ~[]T, R any](slice A, initial R, reducer func(R, T) R) R {
	accum := initial
	for i := len(slice) - 1; i >= 0; i-- {
		accum = reducer(accum, slice[i])
	}
	return accum
}

// ForEach applies fun to each element in slice
//
// fun should be error-safe and handle errors internally. If other arguments are required by predicate,
// predicate should be made a closure with the appropriate variables referenced.
func ForEach[T any, A ~[]T](slice A, fun func(T)) {
	for _, v := range slice {
		fun(v)
	}
}

// Count consumes a slice of a generic type and counts the elements that return true after applying
// the predicate.
//
// Vacuously, empty slices return 0 regardless of the predicate.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func Count[T any, A ~[]T](slice A, predicate func(T) bool) int {
	res := 0
	for _, v := range slice {
		if predicate(v) {
			res++
		}
	}
	return res
}
