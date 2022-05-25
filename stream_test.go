package functools

import (
	"testing"
)

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	iter := ParallelStream(SliceIter(slice), 5)
	s1 := Map(func(e int) int { return e * -1 }, iter)
	s2 := Sort(func(x int, y int) bool { return x < y }, s1)
	ForEach(func(e int) { print(e) }, s2)
}
