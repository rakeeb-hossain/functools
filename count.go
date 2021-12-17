package functools

// Count consumes a slice of a generic type and counts the elements that return true after applying
// the predicate.
//
// Vacuously, empty slices return 0 regardless of the predicate.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func Count[T any](slice []T, predicate func(T) bool) int {
	res := 0
	for _, v := range slice {
		if predicate(v) {
			res++
		}
	}
	return res
}
