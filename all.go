package functools

// All consumes a slice of a generic type and applies the predicate to each element in the slice.
// All return true if and only if no element returns false after applying the predicate.
//
// Vacuously, empty slices return true regardless of the predicate.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}
