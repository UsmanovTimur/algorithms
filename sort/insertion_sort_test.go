package main

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	insertionSort(arr)
	etalon := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		if arr[i] != etalon[i] {
			t.Errorf("out_arr = %d;", arr)
			return
		}
	}
}
