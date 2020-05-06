package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{9, 9, 8, 7, 7, 6, 6, 5, 5, 4, 3, 2, 1, 10}
	index := linearSearch(arr, 5)
	if index != 8 {
		t.Errorf("index = %d; want 8", index)
	}
}
