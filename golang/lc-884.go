package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	set1, set2 := make(map[string]int), make(map[string]int)
	list1, list2 := strings.Split(s1, " "), strings.Split(s2, " ")
	for _, word := range list1 {
		set1[word]++
	}
	for _, word := range list2 {
		set2[word]++
	}
	res := []string{}
	for key, val := range set1 {
		if val == 1 && set2[key] == 0 {
			res = append(res, key)
		}
	}
	for key, val := range set2 {
		if val == 1 && set1[key] == 0 {
			res = append(res, key)
		}
	}
	return res
}
