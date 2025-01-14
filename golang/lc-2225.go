package main

import "sort"

func findWinners(matches [][]int) [][]int {
	table := make(map[int]int)
	for _, match := range matches {
		table[match[1]]++
		if _, ok := table[match[0]]; ok == false {
			table[match[0]] = 0
		}
	}
	ans := make([][]int, 2)
	for key, val := range table {
		if val == 1 {
			ans[1] = append(ans[1], key)
		} else if val == 0 {
			ans[0] = append(ans[0], key)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
