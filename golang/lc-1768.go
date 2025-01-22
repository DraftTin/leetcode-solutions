package main

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	res := make([]byte, 0, len(word1)+len(word2))
	for i < len(word1) && j < len(word2) {
		res = append(res, word1[i], word2[j])
		i++
		j++
	}
	res = append(res, []byte(word1[i:])...)
	res = append(res, []byte(word2[j:])...)
	return string(res)
}
