package functools

func Filter[T any](slice []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
