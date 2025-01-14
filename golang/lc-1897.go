func makeEqual(words []string) bool {
	charMap := make(map[rune]int)
	for _, word := range words {
		for _, c := range word {
			charMap[c]++
		}
	}
	n := len(words)
	for _, count := range charMap {
		if count%n != 0 {
			return false
		}
	}
	return true
}
