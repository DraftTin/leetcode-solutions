package main

func getNext(s string) []int {
	next := make([]int, len(s))
	k, i := -1, 0
	for i < len(s) {
		if k == -1 || s[i] == s[k] {
			i++
			k++
			next[i] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	i, j := 0, 0
	next := getNext(needle)
	for i < len(haystack) && j < len(needle) {
		if j == -1 || needle[j] == haystack[i] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle)
	return 
}
