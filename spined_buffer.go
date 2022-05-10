package functools

var FirstBuffPower int = 4
var MinSpineSize int = 2

type SpinedBuffer[T any] struct {
	// Spine data structures
	// We optimistically assume everything will fit into currBuff in most cases
	currBuff []T
	spines   [][]T

	// Spine state management
	spineIdx          int
	flatIdx           int
	sizePower         int
	capacity          int
	sizeOfPrevBuffers int
}

// Checks if copy is required and copies currBuff to spines
func (s *SpinedBuffer[T]) inflateSpine() {
	if s.spineIdx == 0 && len(s.currBuff) == s.capacity {
		// Create spines
		s.spines = make([][]T, MinSpineSize)

		// Assign currBuff to first spine and set sizeOfPrevBuffers
		s.spines[0] = s.currBuff
		s.spineIdx++
		s.sizeOfPrevBuffers = s.capacity

		// Update subsequent spines
		for i := 1; i < MinSpineSize; i++ {
			s.sizePower++
			s.spines[i] = make([]T, 1<<s.sizePower)
			s.capacity += 1 << s.sizePower
		}
	}
}

func CreateSpinedBuffer[T any]() (s SpinedBuffer[T]) {
	s.spineIdx = 0
	s.flatIdx = 0
	s.sizePower = FirstBuffPower
	s.capacity = 1 << FirstBuffPower

	s.currBuff = make([]T, s.capacity)
	s.spines = nil
}

func (s SpinedBuffer[T]) Len() int {
	return s.flatIdx
}

func (s SpinedBuffer[T]) Capacity() int {
	return s.capacity
}

func (s *SpinedBuffer[T]) Append(elem T) {
	if s.flatIdx < cap(s.currBuff) {
		// Assign elem to currBuff
		s.currBuff[s.flatIdx] = elem
		s.flatIdx++
	} else {
		s.inflateSpine()
		if s.flatIdx == s.capacity {
			// Need to extend capacity

			// Allocate new array into spines, update capacity, and increment spineIdx
			s.sizePower++
			newBuff := make([]T, 1<<s.sizePower)
			s.capacity += 1 << s.sizePower
			// NOTE: this is where the main optimization happens; only need to copy over existing slice
			// pointers, NOT their respective entries
			s.spines = append(s.spines, newBuff)
			s.spineIdx++

			// Assign value to new spine
			s.spines[s.spineIdx][0] = elem
			s.flatIdx++
		} else {
			// Calculate offset to add elem
			offset := s.flatIdx - s.sizeOfPrevBuffers
			s.spines[s.spineIdx][offset] = elem

			s.flatIdx++
		}
	}
}

func (s SpinedBuffer[T]) Flatten() []T {
	if s.spineIdx == 0 {
		return s.currBuff
	}

	res := make([]T, s.flatIdx)
	currIdx := 0
	for i := 0; i < s.spineIdx; i++ {
		for j := 0; j < len(s.spines[i]); j++ {
			res[currIdx] = s.spines[i][j]
		}
	}

	return res
}
