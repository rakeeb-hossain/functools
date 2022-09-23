package functools

import (
	"fmt"
)

const FirstBuffPower int = 4
const MinSpineSize int = 2 // must be >= 1
const SpineExtendCount int = 1

type AbstractBuffer[T any] interface {
	Push(T)
	At(int)
	Flatten()
	Len() int
	Capacity() int
}

// SpinedBuffer is an optimization on a regular slice that doesn't require copying elements on re-sizing.
// This has good performance in cases where an unknown size stream is being processed, since copying from
// re-sizing is minimized.
type SpinedBuffer[T any] struct {
	// Spine data structures
	// We optimistically assume everything will fit into currBuff in most cases
	currBuff []T
	spines   [][]T

	// Spine state management
	sizeOfPrevBuffers []int
	spineIdx          int
	flatIdx           int
	sizePower         int
	capacity          int
	inflated          bool
}

// Checks if copy is required and copies currBuff to spines
func (s *SpinedBuffer[T]) inflateSpine() {
	if s.spineIdx == 0 && s.flatIdx == s.capacity {
		// Create spines
		s.spines = make([][]T, MinSpineSize)

		// Assign currBuff to first spine and set sizeOfPrevBuffers
		s.spines[0] = s.currBuff[:] // should be O(1) since just copying slice
		s.sizeOfPrevBuffers = make([]int, 1, MinSpineSize)
		s.sizeOfPrevBuffers[0] = s.flatIdx

		// Update subsequent spines
		for i := 1; i < MinSpineSize; i++ {
			s.sizePower++
			s.spines[i] = make([]T, 0, 1<<s.sizePower)
			s.capacity += 1 << s.sizePower
		}
	}
	s.inflated = true
}

func CreateSpinedBuffer[T any]() (s SpinedBuffer[T]) {
	s.spineIdx = 0
	s.flatIdx = 0
	s.sizePower = FirstBuffPower
	s.capacity = 1 << FirstBuffPower

	s.currBuff = make([]T, s.capacity)
	s.spines = nil
	s.sizeOfPrevBuffers = nil

	return s
}

func (s SpinedBuffer[T]) Len() int {
	return s.flatIdx
}

func (s SpinedBuffer[T]) Capacity() int {
	return s.capacity
}

func (s *SpinedBuffer[T]) Push(elem T) {
	if s.flatIdx < cap(s.currBuff) {
		// Assign elem to currBuff
		s.currBuff[s.flatIdx] = elem
		s.flatIdx++
	} else {
		if !s.inflated {
			s.inflateSpine()
		}
		// Check if len == cap
		if len(s.spines[s.spineIdx]) == (1 << (s.spineIdx + FirstBuffPower)) {
			// Check if we need to extend capacity
			if s.flatIdx == s.capacity {
				// Allocate new array into spines, update capacity, and increment spineIdx
				for i := 0; i < SpineExtendCount; i++ {
					s.sizePower++
					newBuff := make([]T, 0, 1<<s.sizePower)
					s.capacity += 1 << s.sizePower

					// NOTE: this is where the main optimization happens; only need to copy over existing slice
					// pointers, NOT their respective entries
					s.spines = append(s.spines, newBuff)
				}
			}
			s.spineIdx++

			// Assign value to new spine
			s.spines[s.spineIdx] = append(s.spines[s.spineIdx], elem)
			s.flatIdx++

			// Create new sizeOfPrevBuffers entry
			s.sizeOfPrevBuffers = append(s.sizeOfPrevBuffers, 0)
		} else {
			s.spines[s.spineIdx] = append(s.spines[s.spineIdx], elem)
			s.flatIdx++
		}
		s.sizeOfPrevBuffers[len(s.sizeOfPrevBuffers)-1] = s.flatIdx
	}
}

func (s SpinedBuffer[T]) Flatten() []T {
	if s.spineIdx == 0 {
		return s.currBuff[:s.flatIdx]
	}

	res := make([]T, s.flatIdx)
	currIdx := 0
	for i := 0; i <= s.spineIdx; i++ {
		j := 0
		for ; currIdx < s.sizeOfPrevBuffers[i]; currIdx++ {
			res[currIdx] = s.spines[i][j]
			j++
		}
	}

	return res
}

func (s SpinedBuffer[T]) At(index int) (res T) {
	if index < s.flatIdx && index >= 0 {
		if s.spineIdx == 0 {
			res = s.currBuff[index]
		} else {
			// binary-search for upper-bound; gives index of first elem in sizeOfPrevBuffers that is >= index
			// this index is guaranteed to be valid since sizeOfPrevBuffers last elem is s.flatIdx and index < s.flatIdx
			spineSizeIdx := upperBoundGuaranteed(index, s.sizeOfPrevBuffers)

			// Equality-case where index actually belongs to next spine
			if s.sizeOfPrevBuffers[spineSizeIdx] == index {
				res = s.spines[spineSizeIdx+1][0]
			} else {
				offset := index // case where index belongs to first spine so spineSizeIdx == 0
				if spineSizeIdx > 0 {
					offset = index - s.sizeOfPrevBuffers[spineSizeIdx-1]
				}
				res = s.spines[spineSizeIdx][offset]
			}

		}
	}
	return res
}

func (s SpinedBuffer[T]) PrintStats() {
	fmt.Printf("spineIdx: %d\n", s.spineIdx)
	fmt.Printf("flatIdx: %d\n", s.flatIdx)
	fmt.Printf("capacity: %d\n", s.capacity)
	fmt.Printf("sizePower: %d\n", s.sizePower)
}

func upperBoundGuaranteed(val int, arr []int) int {
	lo := 0
	hi := len(arr)
	for lo < hi {
		mid := (hi-lo)/2 + lo

		if arr[mid] >= val {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
