func isAcronym(words []string, s string) bool {
	if len(words) != s {
		return false
	}
	i := 0
	for _, word := range words {
		if word[0] != s[i] {
			return false
		}
		i++
	}
	return true
}
