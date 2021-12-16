package functools

func Count[T any] (slice []T, predicate func (T) bool) int {
	res := 0
	for _, v := range slice {
		if predicate(v) {
			res++
		}
	}
	return res
}