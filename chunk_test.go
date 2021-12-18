package functools

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	chunk := Chunk(slice, 2)
	if !reflect.DeepEqual(chunk, [][]int{{1, 2}, {3, 4}}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{{1, 2}, {3, 4}}, chunk)
	}

	chunk = Chunk(slice, 10)
	if !reflect.DeepEqual(chunk, [][]int{{1, 2, 3, 4}}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{{1, 2, 3, 4}}, chunk)
	}

	chunk = Chunk(slice, 0)
	if !reflect.DeepEqual(chunk, [][]int{}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{}, chunk)
	}

	chunk = Chunk(slice, 3)
	if !reflect.DeepEqual(chunk, [][]int{{1, 2, 3}, {4}}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{{1, 2, 3}, {4}}, chunk)
	}

	chunk = Chunk([]int{}, -1)
	if !reflect.DeepEqual(chunk, [][]int{}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{}, chunk)
	}

	chunk = Chunk(slice, 1)
	if !reflect.DeepEqual(chunk, [][]int{{1}, {2}, {3}, {4}}) {
		t.Errorf("TestChunk failed, expected %v, but got %v", [][]int{{1}, {2}, {3}, {4}}, chunk)
	}
}
