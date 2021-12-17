package functools

// Reduce consumes a slice of a generic type and an initial value. It reduces the slice to a single value by applying
// the binary reducer function to each successive element in the slice.
//
// Vacuously, empty slices return the initial value provided.
//
// reducer should be error-safe. It should handle any errors internally and return the desired type.
// If other arguments are required by reducer, reducer should be made a closure with the appropriate
// variables referenced.
func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
	accum := initial
	for _, v := range slice {
		accum = reducer(accum, v)
	}
	return accum
}
