package main

import "fmt"

// Сортировка вставками
func main() {
	arr := []int{9, 9, 8, 7, 7, 6, 6, 5, 5, 4, 3, 2, 1, 10}
	insertionSort(arr)
	fmt.Printf("%v\n", arr)
}
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i - 1
		value := a[i]
		for j >= 0 && a[j] > value {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = value
	}
}
