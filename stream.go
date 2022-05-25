package functools

const (
	SIZED = 1 << iota
	SORTED
	DISTINCT
	ORDERED
	UNORDERED
)

type StreamStage[T any] interface {
	spliterator() Spliterator[T]

	getParallelism() int
	isStateful() bool
	// characteristics() uint32

	opEvalParallelLazy()
}

// StatelessOp struct embedding
type InheritUpstream[T any] struct {
	upstream *StreamStage[T]
}

func (s InheritUpstream[T]) getParallelism() int {
	return (*s.upstream).getParallelism()
}

type StatelessOp struct{}

func (s StatelessOp) isStateful() bool {
	return false
}

func (s StatelessOp) opEvalParallelLazy() {}

// SourceStage definition
type SourceStage[T any] struct {
	StatelessOp
	src         Spliterator[T]
	parallelism int
}

func Stream[T any](spliterator Spliterator[T]) StreamStage[T] {
	return SourceStage[T]{src: spliterator}
}

func ParallelStream[T any](spliterator Spliterator[T], parallelism int) StreamStage[T] {
	return SourceStage[T]{src: spliterator, parallelism: parallelism}
}

func (s SourceStage[T]) spliterator() Spliterator[T] {
	return s.src
}

func (s SourceStage[T]) getParallelism() int {
	return s.parallelism
}

// All this stuff should probably go into a separate file

// Helpers
func UpstreamToBuffer[T any](src StreamStage[T]) []T {
	slice := make([]T, 0)
	src.spliterator().forEachRemaining(func(e T) {
		slice = append(slice, e)
	})
	return slice
}

// Map
type MapOp[TIn any, TOut any] struct {
	StatelessOp
	InheritUpstream[TIn]
	mapper func(TIn) TOut
}

func Map[TIn any, TOut any](mapper func(TIn) TOut, upstream StreamStage[TIn]) StreamStage[TOut] {
	return MapOp[TIn, TOut]{
		StatelessOp{},
		InheritUpstream[TIn]{upstream: &upstream},
		mapper,
	}
}

func mapSpliterator[T any, O any](mapper func(T) O, src Spliterator[T]) (res Spliterator[O]) {
	res.tryAdvance = func(fn func(O)) bool {
		wrapper_fn := func(e T) {
			v := mapper(e)
			fn(v)
		}
		return src.tryAdvance(wrapper_fn)
	}
	res.forEachRemaining = func(fn func(O)) {
		wrapper_fn := func(e T) {
			v := mapper(e)
			fn(v)
		}
		src.forEachRemaining(wrapper_fn)
	}
	// Recursive split!!!
	res.trySplit = func() (Spliterator[O], bool) {
		r, b := src.trySplit()
		if !b {
			return Spliterator[O]{}, false
		} else {
			return mapSpliterator[T, O](mapper, r), true
		}
	}
	return res
}

func (m MapOp[TIn, TOut]) spliterator() (res Spliterator[TOut]) {
	s := (*m.InheritUpstream.upstream).spliterator()
	return mapSpliterator[TIn, TOut](m.mapper, s)
}

// SortOp
type SortOp[T any] struct {
	InheritUpstream[T]
	cmp func(T, T) bool
}

func Sort[T any](cmp func(T, T) bool, upstream StreamStage[T]) StreamStage[T] {
	return SortOp[T]{
		InheritUpstream[T]{upstream: &upstream},
		cmp,
	}
}

func quicksort[T any](cmp func(T, T) bool, slice []T) {
	for i, _ := range slice {
		min_so_far := slice[i]
		min_ind := i
		for j := i + 1; j < len(slice); j++ {
			if cmp(slice[j], min_so_far) {
				min_so_far = slice[j]
				min_ind = j
			}
		}
		tmp := slice[i]
		slice[i] = slice[min_ind]
		slice[min_ind] = tmp
	}
}

func (m SortOp[T]) spliterator() (res Spliterator[T]) {
	s := (*m.upstream).spliterator()
	done := false
	buffer := make([]T, 0, 2)

	res.tryAdvance = func(fn func(T)) bool {
		if !done {
			s.forEachRemaining(func(e T) {
				buffer = append(buffer, e)
			})
			quicksort(m.cmp, buffer)
			done = true
		}
		fn(buffer[0])
		return true
	}
	res.forEachRemaining = func(fn func(T)) {
		if !done {
			s.forEachRemaining(func(e T) {
				buffer = append(buffer, e)
			})
			quicksort(m.cmp, buffer)
			done = true
		}
		for _, x := range buffer {
			fn(x)
		}
	}
	res.trySplit = s.trySplit
	return res
}

func (s SortOp[T]) isStateful() bool {
	return true
}

func (s SortOp[T]) opEvalParallelLazy() {

}

// Terminal ops
func ForEach[T any](fn func(T), stream StreamStage[T]) {
	n := stream.getParallelism()
	if n <= 1 {
		stream.spliterator().forEachRemaining(fn)
	} else {
		// FILL IN
	}
}

func test() {
	slice := make([]int, 5)
	s := Stream(SliceIter(slice))
	ForEach(func(n int) { print(n) }, s)
}
