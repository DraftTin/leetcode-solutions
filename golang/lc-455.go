package main

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i := len(g) - 1
	for g[i] > s[len(s)-1] {
		i--
	}
	j := len(s) - 1
	ans := 0
	for i >= 0 && j >= 0 {
		if g[i] <= s[j] {
			ans++
			j--
			i--
		} else {
			i--
		}
	}
	return ans
}
