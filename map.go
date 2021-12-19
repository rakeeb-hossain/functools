package functools

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