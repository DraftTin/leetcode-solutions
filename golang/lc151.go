package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(strings.Trim(s, " "), " ")

	reversed := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) == 0 {
			continue
		}
		reversed = append(reversed, words[i])
	}
	return strings.Join(reversed, " ")
}
