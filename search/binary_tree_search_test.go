package search

import (
	"testing"
)

func TestBinaryTreeSearch(t *testing.T) {
	arr := []int{6, 5, 4, 8, 3, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	index := BinaryTreeSearch(arr, 7)
	if index != 12 {
		t.Errorf("index = %d; want 7", index)
	}
}
