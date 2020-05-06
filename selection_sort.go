package main

import "fmt"

//Сортировка выбором
func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	arr = selectionSort(arr)
	fmt.Printf("%v\n", arr)
}
func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		// Ищем минимальный элемент
		min_elem := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min_elem] {
				min_elem = j
			}
		}
		if min_elem != i {
			a[i] = a[i] + a[min_elem]
			a[min_elem] = a[i] - a[min_elem]
			a[i] = a[i] - a[min_elem]
		}
	}
	return a
}
