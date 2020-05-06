package main

// Сортировка вставками
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
