package functools

type Spliterator[T any] interface {
	TryAdvance(func(T)) bool
	ForEachRemaining(func(T))
	TrySplit() (Spliterator[T], bool)
}

type sliceSpliterator[T any] struct {
	Slice 		 []T
	SliceLength  int
	index		 int
}

func Iter[T any, A ~[]T](slice A) sliceSpliterator[T] {
	res := sliceSpliterator[T]{}
	res.Slice = slice
	res.SliceLength = len(slice)
	return res
}

func (ss sliceSpliterator[T]) TryAdvance(fn func(T)) bool {
	if ss.index >= ss.SliceLength {
		return false // b is false here
	}
	ss.index++
	fn(ss.Slice[ss.index-1])
	return true
}

func (ss sliceSpliterator[T]) ForEachRemaining(fn func(T)) {
	for ; ss.index < ss.SliceLength ; ss.index++ {
		fn(ss.Slice[ss.index])
	}
}

func (ss *sliceSpliterator[T]) TrySplit() (res sliceSpliterator[T], b bool) {
	if ss.SliceLength <= 1 {
		return res, b
	}
	updatedLen := ss.SliceLength / 2
	res.SliceLength = ss.SliceLength - updatedLen
	res.Slice = ss.Slice[updatedLen:]
	ss.SliceLength = updatedLen

	return res, true
}



type ruleSpliterator[T any] struct {
	Rule 	func() (T, bool)
}

func RuleIter[T any](rule func() (T, bool)) ruleSpliterator[T] {
	return ruleSpliterator[T]{Rule: rule}
}

func (rs ruleSpliterator[T]) TryAdvance(fn func(T)) bool {
	res, ok := rs.Rule()
	if ok {
		fn(res)
	}
	return ok
}

func (rs ruleSpliterator[T]) ForEachRemaining(fn func(T)) {
	for res, ok := rs.Rule(); ok; res, ok = rs.Rule() {
		fn(res)
	}
}

func (rs ruleSpliterator[T]) TrySplit() (res ruleSpliterator[T], b bool) {
	return res, b
}



type intRangeSpliterator struct {
	low int
	high int
}

func RangeIter(low, high int) intRangeSpliterator {
	return intRangeSpliterator{low: low, high: high}
}

func (is intRangeSpliterator) TryAdvance(fn func(int)) bool {
	if is.low >= is.high {
		return false
	}
	is.low++
	fn(is.low-1)
	return true
}

func (is intRangeSpliterator) ForEachRemaining(fn func(int)) {
	for ; is.low < is.high; is.low++ {
		fn(is.low)
	}
}

func (is *intRangeSpliterator) TrySplit() (res intRangeSpliterator, b bool) {
	if is.high - is.low <= 1 {
		return res, b
	}
	avg := (is.low + is.high)/2
	is.high = avg
	res.low = avg
	return res, true
}



// Iterator is a generic iterator on a slice that lazily evaluates the next element in the slice.
// This is used to lazily evaluate a slice's next value, allowing several applications of functional
// methods on a single list while only incurring a O(1) memory overhead.
type Iterator[T any] func() (T, bool)

// Iter consumes a generic slice and generates a forward-advancing Iterator
//
// Iter is passed a copy of the slice. This does not copy the contents of the slice, but the size of
// the slice is fixed. Therefore, modifications of element the internal slice will affect the Iterator
func Iter[T any, A ~[]T](slice A) Iterator[T] {
	index := 0
	return func() (t T, b bool) {
		if index >= len(slice) {
			return t, b // b is false here
		}
		index++
		return slice[index-1], true
	}
}

// ReverseIter consumes a generic slice and generates a reverse-advancing Iterator
func ReverseIter[T any, A ~[]T](slice A) Iterator[T] {
	index := len(slice) - 1
	return func() (t T, b bool) {
		if index < 0 || index >= len(slice) {
			return t, b // b is false here
		}
		index--
		return slice[index+1], true
	}
}

// Slice converts a generic Iterator to a slice of the appropriate type
func Slice[T any](iter Iterator[T]) []T {
	res := make([]T, 0)
	for val, ok := iter(); ok; val, ok = iter() {
		res = append(res, val)
	}
	return res
}

// Next is an alias for advancing the Iterator
func Next[T any](iter Iterator[T]) (T, bool) {
	return iter()
}
