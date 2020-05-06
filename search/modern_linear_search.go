package search

//Модернизированный линейный поиск

func modernLinearSearch(arr []int, find_elem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == find_elem {
			return i
		}
	}
	return -1
}
