package functools

func Map[T any, R any] (slice []T, mapper func(T) R) []R {
	res := make([]R, len(slice))

	for i, v := range slice {
		res[i] = mapper(v)
	}
	return res
}