package main

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	pos := []int{}
	for ind, _ := range s {
		if vowels[s[ind]] == true {
			pos = append(pos, ind)
		}
	}
	i, j := 0, len(pos)-1
	res := []byte(s)
	for i < j {
		res[pos[i]], res[pos[j]] = res[pos[j]], res[pos[i]]
		i++
		j--
	}
	return string(res)
}
