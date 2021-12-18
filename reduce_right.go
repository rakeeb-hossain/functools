package functools

// ReduceRight consumes a slice of a generic type and an initial value. It
// reduces the slice to a single value by applying the binary reducer function
// to each element in the slice. ReduceRight differs from Reduce by iterating
// from the last element to the first element.
//
// Vacuously, empty slices return the initial value provided.
//
// reducer should error-safe. It should handle any errors internaly and return
// the desired type. If other arguments are required by reducer, reducer should
// be made a closure with the appropriate variables referenced.
func ReduceRight[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
	accum := initial
	for i := len(slice) - 1; i >= 0; i-- {
		accum = reducer(accum, slice[i])
	}
	return accum
}
