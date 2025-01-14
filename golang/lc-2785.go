package main

func sortVowels(s string) string {
	vowels := map[byte]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0, 'A': 0, 'E': 0, 'I': 0, 'O': 0, 'U': 0}
	pos := []int{}
	for i := 0; i < len(s); i++ {
		if _, ok := vowels[s[i]]; ok == true {
			vowels[s[i]]++
			pos = append(pos, i)
		}
	}
	res := []byte(s)
	k := 0
	tmp := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	for _, c := range tmp {
		for i := 0; i < vowels[c]; i++ {
			res[pos[k]] = c
			k++
		}
	}
	return string(res)
}
