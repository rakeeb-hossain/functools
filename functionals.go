// Contains classic generic functional methods

package functools

import "math"

// Map consumes a slice of a generic type and returns a slice with the supplied mapping function
// applied to each element.
//
// mapper should be error-safe. It should handle any errors internally and return the desired type.
// If other arguments are required by mapper, mapper should be made a closure with the appropriate
// variables referenced.
func Map[T any, A ~[]T, R any](slice A, mapper func(T) R) []R {
	res := make([]R, len(slice))

	for i, v := range slice {
		res[i] = mapper(v)
	}
	return res
}

func MapIter[T any, R any](iter Iterator[T], mapper func(T) R) Iterator[R] {
	return func() (R, bool) {
		res, ok := Next(iter)
		return mapper(res), ok
	}
}

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
func Reduce[T any, A ~[]T, R any](slice A, initial R, reducer func(R, T) R) R {
	accum := initial
	for _, v := range slice {
		accum = reducer(accum, v)
	}
	return accum
}

func ReduceIter[T any, R any](iter Iterator[T], initial R, reducer func(R, T) R) R {
	accum := initial
	for val, ok := Next(iter); ok; val, ok = Next(iter) {
		accum = reducer(accum, val)
	}
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

// Chunk consumes a slice of a generic type and
// returns a slice composed of multiple chunks of a user-specified size
//
// Vacuously, empty slices return an empty slice of slices regardless of the chunk size.
//
// When the chunk size is zero or negative, return an empty slice
func Chunk[T any, A ~[]T](slice A, size int) []A {
	n := len(slice)
	if size <= 0 || n == 0 {
		return []A{}
	}

	if size >= n {
		tmp := make(A, n)
		copy(tmp, slice)
		return []A{tmp}
	}

	res := make([]A, 0, int(math.Ceil(float64(n/size))))
	for i := 0; i < n; i += size {
		if i+size <= n {
			tmp := make(A, size)
			copy(tmp, slice[i:i+size])
			res = append(res, tmp)
		} else {
			tmp := make(A, n-i)
			copy(tmp, slice[i:])
			res = append(res, tmp)
		}
	}
	return res
}
