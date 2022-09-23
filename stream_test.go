package functools

import (
	"testing"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestAny(t *testing.T) {
	slice := make([]int, 10000000)
	for i, _ := range slice {
		slice[i] = i
	}
	iter := ParallelStream(SliceIter(slice), 5)
	// s2 := Sort(func(x int, y int) bool { return x < y }, s1)

	print(Any(isPrime, iter))
}

func BenchmarkMap(b *testing.B) {
	slice := make([]int, 10000000)
	for i, _ := range slice {
		slice[i] = i
	}
	iter := ParallelStream(SliceIter(slice), 100)
	s1 := Map(func(e int) int { return e * -1 }, iter)
	// s2 := Sort(func(x int, y int) bool { return x < y }, s1)
	Sum(s1)
}

func BenchmarkMapSeq(b *testing.B) {
	slice := make([]int, 10000000)
	for i, _ := range slice {
		slice[i] = i
	}
	iter := Stream(SliceIter(slice))
	s1 := Map(func(e int) int { return e * -1 }, iter)
	// s2 := Sort(func(x int, y int) bool { return x < y }, s1)
	Sum(s1)
}

func BenchmarkMapFor(b *testing.B) {
	slice := make([]int, 10000000)
	for i, _ := range slice {
		slice[i] = i
	}
	for i, _ := range slice {
		slice[i] *= -1
	}
	res := 0
	for _, v := range slice {
		res += v
	}
}
