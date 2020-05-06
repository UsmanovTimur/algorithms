package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	index := binarySearch(arr, 7)
	if index != 7 {
		t.Errorf("index = %d; want 7", index)
	}
}
