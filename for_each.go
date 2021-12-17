package functools

// ForEach applies fun to each element in slice
//
// fun should be error-safe and handle errors internally. If other arguments are required by predicate,
// predicate should be made a closure with the appropriate variables referenced.
func ForEach[T any] (slice []T, fun func(T)) {
	for _, v := range slice {
		fun(v)
	}
}