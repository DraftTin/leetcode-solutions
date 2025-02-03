package main

import "sort"

func minimumKeypresses(s string) int {
	counter := make([]int, 27)
	for _, c := range s {
		counter[c-'a']++
	}
	sort.Ints(counter)
	res := 0
	keyNum := 9
	pressNum := 1
	for i := len(counter) - 1; i >= 0; i-- {
		if counter[i] == 0 {
			break
		}
		res += pressNum * counter[i]
		keyNum--
		if keyNum == 0 {
			pressNum++
			keyNum = 9
		}
	}
	return res
}
