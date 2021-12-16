package functools

func Reduce[T any, R any] (slice []T, initial R, reducer func(R, T) R) R {
	accum := initial
	for _, v := range slice {
		accum = reducer(accum, v)
	}
	return accum
}