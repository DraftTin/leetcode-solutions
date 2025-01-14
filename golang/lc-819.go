package main

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	curStr := ""
	paragraph = strings.ToLower(paragraph)
	wordCounts := make(map[string]int)
	for _, c := range paragraph {
		if strings.Contains("!?;,'. ", string(c)) {
			wordCounts[curStr]++
			curStr = ""
		} else {
			curStr += string(c)
		}
	}
	for _, word := range banned {
		delete(wordCounts, word)
	}
	maxCount, ans := 0, ""
	for word, count := range wordCounts {
		if count > maxCount {
			maxCount = count
			ans = word
		}
	}
	return ans
}
