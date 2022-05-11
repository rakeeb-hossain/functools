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

const BENCHMARK_SIZE int = 10000

func BenchmarkCreateSpinedBuffer_Push(b *testing.B) {
	buff := CreateSpinedBuffer[int]()

	for i := 0; i < BENCHMARK_SIZE; i++ {
		buff.Push(i)
	}
}

func BenchmarkSlicePush(b *testing.B) {
	slice := make([]int, 0)

	for i := 0; i < BENCHMARK_SIZE; i++ {
		slice = append(slice, i)
	}
}

//func BenchmarkArrayPush(b *testing.B) {
//	arr := [BENCHMARK_SIZE]int{}
//
//	for i := 0; i < BENCHMARK_SIZE; i++ {
//		arr[i] = i
//	}
//}
