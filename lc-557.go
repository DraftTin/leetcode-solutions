package main

import "strings"

func reverseWords(s string) string {
	splitS := strings.Split(s, " ")
	ans := make([]string, 0, len(splitS))
	for _, s := range splitS {
		reversed := []byte(s)
		i, j := 0, len(reversed)-1
		for i < j {
			reversed[i], reversed[j] = reversed[j], reversed[i]
			i++
			j--
		}
		ans = append(ans, string(reversed))
	}
	return strings.Join(ans, " ")
}
