package functools

// All consumes a slice of a generic type and applies the predicate to each element in the slice.
// All return true if and only if no element returns false after applying the predicate.
//
// Vacuously, empty slices return true regardless of the predicate.
//
// predicate should be error-safe. It should handle any errors internally and return only a bool.
// If other arguments are required by predicate, predicate should be made a closure with the appropriate
// variables referenced.
func All[T any, A ~[]T](slice A, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

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

// Summable encompasses all builtin types with the + operator defined on them or any type aliases
// of these types
type Summable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Sum consumes a slice of a Summable type and sums the elements
//
// Vacuously, empty slices return the zero value of the provided Summable
func Sum[S Summable, A ~[]S](slice A) S {
	var res S
	for _, v := range slice {
		res += v
	}
	return res
}
