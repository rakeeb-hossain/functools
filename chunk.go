package functools

import (
	"math"
)

// Chunk consumes a slice of a generic type and
// returns a slice composed of multiple chunks of a user-specified size
//
// Vacuously, empty slices return an empty slice of slices regardless of the chunk size.
//
// When the chunk size is zero or negative, return an empty slice
func Chunk[T any, A ~[]T](slice A, size int) []A {
	n := len(slice)
	if size <= 0 || n == 0 {
		return []A{}
	}

	if size >= n {
		tmp := make(A, n)
		copy(tmp, slice)
		return []A{tmp}
	}

	res := make([]A, 0, int(math.Ceil(float64(n/size))))
	for i := 0; i < n; i += size {
		if i+size <= n {
			tmp := make(A, size)
			copy(tmp, slice[i:i+size])
			res = append(res, tmp)
		} else {
			tmp := make(A, n-i)
			copy(tmp, slice[i:])
			res = append(res, tmp)
		}
	}
	return res
}
