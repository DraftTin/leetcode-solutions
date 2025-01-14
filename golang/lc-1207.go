package main

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, val := range arr {
		counter[val] += 1
	}
	table := make(map[int]bool)
	for _, val := range counter {
		if table[val] == true {
			return false
		}
		table[val] = true
	}
	return true
}
