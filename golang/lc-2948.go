package main

import (
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	clone := make([]int, len(nums))
	copy(clone, nums)
	sort.Ints(clone)
	groups := [][]int{[]int{clone[0]}}
	groupMark := make(map[int]int)
	groupMark[clone[0]] = 0
	for i := 1; i < len(nums); i++ {
		if clone[i]-clone[i-1] > limit {
			groups = append(groups, []int{})
		}
		groupMark[clone[i]] = len(groups) - 1
		groups[len(groups)-1] = append(groups[len(groups)-1], clone[i])
	}
	for i := 0; i < len(nums); i++ {
		groupNum := groupMark[nums[i]]
		nums[i] = groups[groupNum][0]
		groups[groupNum] = groups[groupNum][1:]
	}
	return nums
}
