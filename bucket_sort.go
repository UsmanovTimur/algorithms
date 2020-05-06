package main

import "fmt"

// Блочная сортировка
func main() {
	array := []int{4, 1, 1, 3, 2, 2, 7, -6, 0}
	fmt.Printf("Before Bucket Sort %v\n", array)
	array = bucketSort(array, 5)
	fmt.Printf("After Bucket Sort %v\n", array)
}

func bucketSort(array []int, bucketSize int) []int {
	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			__insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func __insertionSort(a []int) {
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
