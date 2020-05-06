package search

//Бинарный поиск

func binarySearch(arr []int, find_elem int) int {
	low := 0
	high := len(arr)
	for low <= high {
		mid := int((high + low) / 2)
		if find_elem == arr[mid] {
			return mid
		}
		if find_elem > arr[mid] {
			low = mid
		}
		if find_elem < arr[mid] {
			high = mid
		}
	}
	return -1
}
