package search

//Последовательный поиск, линейный поиск

func linearSearch(arr []int, find_elem int) int {
	out := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == find_elem {
			out = i
		}
	}
	return out
}
