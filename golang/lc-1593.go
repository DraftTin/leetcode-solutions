package main

// backtracking
func maxUniqueSplit(s string) int {
	return dfs1593(s, 0, map[string]bool{})
}

func dfs1593(s string, curPos int, seen map[string]bool) int {
	if curPos == len(s) {
		return 0
	}
	res := 0
	for i := curPos; i < len(s); i++ {
		curStr := s[curPos : i+1]
		if seen[curStr] {
			continue
		}
		seen[curStr] = true
		res = max(res, dfs1593(s, i+1, seen)+1)
		seen[curStr] = false
	}
	return res
}
