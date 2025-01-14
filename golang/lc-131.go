func partition(s string) [][]string {
	result := [][]string{}
	partition := []string{}
	dfs(s, &partition, &result)
	return result
}

func dfs(s string, partition *[]string, result *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(*partition))
		copy(tmp, *partition)
		*result = append(*result, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			prefix := s[:i]
			postpart := s[i:]
			*partition = append(*partition, prefix)
			dfs(postpart, partition, result)
			*partition = (*partition)[:len(*partition)-1]
		}
	}
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
