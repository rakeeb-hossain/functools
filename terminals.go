// Terminal ops

package functools

import "sync"

// Helpers
func buildNSplits[T any](n uint32, src Spliterator[T]) []Spliterator[T] {
	if n <= 1 {
		return []Spliterator[T]{src}
	}
	// Round N down to a power of 2
	var mask uint32 = 1 << 31
	for n&mask == 0 {
		mask >>= 1
	}
	// Alloc results slice
	res := make([]Spliterator[T], 0, mask)

	// In-order traversal of split tree
	var buildNSplitsRec func(uint32, Spliterator[T])
	buildNSplitsRec = func(n uint32, src Spliterator[T]) {
		if n == 1 {
			res = append(res, src)
		} else {
			split, ok := src.trySplit()
			buildNSplitsRec(n/2, src)
			if ok {
				buildNSplitsRec(n/2, split)
			}
		}
	}
	buildNSplitsRec(mask, src)

	return res
}

// ForEach
func ForEach[T any](fn func(T), stream StreamStage[T]) {
	n := stream.getParallelism()
	if n <= 1 {
		stream.spliterator().forEachRemaining(fn)
	} else {
		// Evaluate up to last stateful op
		stream.opEvalParallelLazy(n)

		// Get n splits
		splits := buildNSplits(uint32(n), stream.spliterator())
		n = len(splits)

		// Perform go-routines
		var wg sync.WaitGroup

		for i := 0; i < n; i++ {
			splits[i].forEachRemaining(fn)

			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				// TODO: abstract into evalSequential so you don't need to rewrite this everytime
				splits[i].forEachRemaining(fn)
			}(i)
		}
		wg.Wait()
	}
}

//// All consumes a slice of a generic type and applies the predicate to each element in the slice.
//// All return true if and only if no element returns false after applying the predicate.
////
//// Vacuously, empty slices return true regardless of the predicate.
////
//// predicate should be error-safe. It should handle any errors internally and return only a bool.
//// If other arguments are required by predicate, predicate should be made a closure with the appropriate
//// variables referenced.
//func All[T any, A ~[]T](slice A, predicate func(T) bool) bool {
//	for _, v := range slice {
//		if !predicate(v) {
//			return false
//		}
//	}
//	return true
//}
//
//// Any consumes a slice of a generic type and applies the predicate to each element in the slice.
//// If any element returns true after applying the predicate, Any returns true.
////
//// Vacuously, empty slices return false regardless of the predicate.
////
//// predicate should be error-safe. It should handle any errors internally and return only a bool.
//// If other arguments are required by predicate, predicate should be made a closure with the appropriate
//// variables referenced.
//func Any[T any, A ~[]T](slice A, predicate func(T) bool) bool {
//	for _, v := range slice {
//		if predicate(v) {
//			return true
//		}
//	}
//	return false
//}
//
//// Summable encompasses all builtin types with the + operator defined on them or any type aliases
//// of these types
//type Summable interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64 |
//		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//		~float32 | ~float64 |
//		~string
//}
//
//// Sum consumes a slice of a Summable type and sums the elements
////
//// Vacuously, empty slices return the zero value of the provided Summable
//func Sum[S Summable, A ~[]S](slice A) S {
//	var res S
//	for _, v := range slice {
//		res += v
//	}
//	return res
//}
