func findspecialinteger(arr []int) int {
	mp := map[int]int{}
	for _, num := range arr {
		mp[num]++
	}
	n := len(arr)
	for key, val := range mp {
		if val > n/4 {
			return key
		}
	}
	return 0
}
