func partitionLabels(s string) []int {
	ans := []int{}
	i := 0
	for i < len(s) {
		size := helper(s, i)
		ans = append(ans, size)
		i += size
	}
	return ans
}

func helper(s string, index int) int {
	curMap := make(map[byte]bool)
	curMap[s[index]] = true
	size := 1
	for i := index + 1; i < len(s); i++ {
		if curMap[s[i]] == true {
			for j := index + size; j < i; j++ {
				curMap[s[j]] = true
			}
			size = i - index + 1
		}
	}
	return size
}
