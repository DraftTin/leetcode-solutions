func leastInterval(tasks []byte, n int) int {
	mp := make(map[byte]int)
	maxCount := 0
	for _, c := range tasks {
		mp[c]++
		maxCount = max(maxCount, mp[c])
	}
	ans := (maxCount - 1) * (n + 1)
	for _, count := range mp {
		if count == maxCount {
			ans++
		}
	}
	return max(ans, len(tasks))
}

// A B C
// A B C
// A B D
// A B #
// A
