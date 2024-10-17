package main

func countConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[rune]bool)
	for _, c := range allowed {
		allowedSet[c] = true
	}
	res := len(words)
	for _, word := range words {
		for _, c := range word {
			if allowedSet[c] == false {
				res--
				break
			}
		}
	}
	return res
}
