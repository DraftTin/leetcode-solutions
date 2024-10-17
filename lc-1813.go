package main

import "strings"

// Greedy
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1, words2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	long, short := words1, words2
	if len(words2) > len(short) {
		short = words1
		long = words2
	}
	left, right := 0, len(short)-1
	for left <= right {
		if short[left] != long[left] {
			break
		}
		left++
	}
	r := len(long) - 1
	for right >= left {
		if short[right] != long[r] {
			break
		}
		r--
		right--
	}
	return left > right
}
