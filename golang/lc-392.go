package main

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
		if len(s)-i > len(t)-j {
			return false
		}
	}
	return i == len(s)
}
