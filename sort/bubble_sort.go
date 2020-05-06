package main

// Сортировка пузырьком
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
