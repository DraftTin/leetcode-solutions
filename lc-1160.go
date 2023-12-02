func countCharacters(words []string, chars string) int {
	mp := make(map[rune]int)
	for _, c := range chars {
		mp[c]++
	}
	ans := 0
	for _, word := range words {
		tmp := map[rune]int{}
		for key, count := range mp {
			tmp[key] = count
		}
		flag := true
		for _, c := range word {
			tmp[c]--
			if tmp[c] < 0 {
				flag = false
				break
			}
		}
		if flag {
			ans += len(word)
		}
	}
	return ans
}
