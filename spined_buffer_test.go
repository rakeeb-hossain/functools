package functools

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSpinedBuffer_Push(t *testing.T) {
	NUM_TRIALS := 10000
	buff := CreateSpinedBuffer[int]()
	arr := make([]int, 0, NUM_TRIALS)

	for i := 1; i <= NUM_TRIALS; i++ {
		buff.Push(i)
		arr = append(arr, i)
		if !reflect.DeepEqual(arr, buff.Flatten()) {
			buff.PrintStats()
			t.Errorf("error at %d", i)
		}
	}
}

func TestSpinedBuffer_At(t *testing.T) {
	NUM_TRIALS := 10000
	buff := CreateSpinedBuffer[int]()

	for i := 1; i <= NUM_TRIALS; i++ {
		buff.Push(i)

		if buff.Len() != i {
			t.Errorf("error at %d", i)
		}

		r := rand.Int() % buff.Len()
		if buff.At(r) != (r + 1) {
			t.Errorf("error at index %d on iteration %d", r, i)
		}
	}
}
