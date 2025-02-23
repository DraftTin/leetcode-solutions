package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
	}
	countList := make([][2]int, len(set))
	for key, val := range set {
		countList = append(countList, [2]int{key, val})
	}
	sort.Slice(countList, func(i, j int) bool {
		return countList[i][1] >= countList[j][1]
	})
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, countList[i][0])
	}
	return res
}
