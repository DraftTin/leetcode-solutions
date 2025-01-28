package main

func partitionString(s string) int {
	curSet := map[byte]bool{}
	i := 0
	res := 0
	for i < len(s) {
		if curSet[s[i]] {
			curSet = map[byte]bool{}
			res++
		}
		curSet[s[i]] = true
	}
	return res + 1
}
