package functools

import (
	"math"
)

// Chunk consumes a slice of a generic type and
// returns a slice composed of multiple chunks of a user-specified size
//
// Vacuously, empty slices return empty regardless of the chunk size.
//
// When the chunk size is zero or negative, return an empty slice
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	n := len(slice)
	if size >= n {
		tmp := make([]T, n)
		copy(tmp, slice)
		return [][]T{tmp}
	}

	res := make([][]T, 0, int(math.Ceil(float64(n/size))))
	for i := 0; i < n; i += size {
		if i+size <= n {
			tmp := make([]T, size)
			copy(tmp, slice[i:i+size])
			res = append(res, tmp)
		} else {
			tmp := make([]T, n-i)
			copy(tmp, slice[i:])
			res = append(res, tmp)
			break
		}
	}
	return res
}
