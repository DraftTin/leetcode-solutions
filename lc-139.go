func wordBreak(s string, wordDict []string) bool {
	visited := make([]bool, len(s))
	return helper(s, 0, wordDict, visited)
}

func helper(s string, i int, wordDict []string, visited []bool) bool {
	if i == len(s) {
		return true
	}
	if visited[i] == true {
		return false
	}
	for _, word := range wordDict {
		if len(word)+i <= len(s) && word == s[i:i+len(word)] {
			if helper(s, i+len(word), wordDict, visited) {
				return true
			}
		}
	}
	visited[i] = true
	return false
}
