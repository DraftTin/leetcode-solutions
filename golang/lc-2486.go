package main

// two pointers problem
func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return len(t) - j
}
