package main

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	out_arr := selectionSort(arr)
	etalon := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(out_arr); i++ {
		if out_arr[i] != etalon[i] {
			t.Errorf("out_arr = %d;", out_arr)
			return
		}
	}
}
