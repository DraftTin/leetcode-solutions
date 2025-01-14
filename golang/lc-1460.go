package main

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	set1, set2 := make(map[int]int), make(map[int]int)
	for _, val := range target {
		set1[val]++
	}
	for _, val := range arr {
		set2[val]++
		if set2[val] > set1[val] {
			return false
		}
	}
	return true
}
