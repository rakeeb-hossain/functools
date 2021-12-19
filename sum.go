package functools

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
