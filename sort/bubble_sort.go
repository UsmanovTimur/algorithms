package main

import "fmt"

// Сортировка пузырьком
func main() {
	arr := []int{9, 9, 8, 7, 7, 6, 6, 5, 5, 4, 3, 2, 1, 10}
	arr = bubleSort(arr)
	fmt.Printf("%v\n", arr)
}
func bubleSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		for i := 0; i < len(a)-j; i++ {
			if a[i] > a[i+1] {
				a[i] = a[i] + a[i+1]
				a[i+1] = a[i] - a[i+1]
				a[i] = a[i] - a[i+1]
			}
		}
	}
	return a
}
