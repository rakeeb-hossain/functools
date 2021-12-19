package functools

// Filter consumes a slice of a generic type and returns a slice with only the elements which returned true after
// applying the predicate.
//
// Elements are returned in the same order that they were supplied in the slice.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func Filter[T any, A ~[]T](slice A, predicate func(T) bool) []T {
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}