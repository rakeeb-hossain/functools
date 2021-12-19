package functools

// Any consumes a slice of a generic type and applies the predicate to each element in the slice.
// If any element returns true after applying the predicate, Any returns true.
//
// Vacuously, empty slices return false regardless of the predicate.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func Any[T any, A ~[]T](slice A, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}
