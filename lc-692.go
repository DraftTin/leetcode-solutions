package main

import "sort"

type Word struct {
	val       string
	frequency int
}

func topKFrequent(words []string, k int) []string {
	wordFrequency := make(map[string]int)
	for _, word := range words {
		wordFrequency[word]++
	}
	wordList := make([]Word, 0, len(wordFrequency))
	for key, val := range wordFrequency {
		wordList = append(wordList, Word{val: key, frequency: val})
	}
	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].frequency == wordList[j].frequency {
			return wordList[i].val <= wordList[j].val
		}
		return wordList[i].frequency >= wordList[j].frequency
	})
	ans := make([]string, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, wordList[i].val)
	}
	return ans
}
