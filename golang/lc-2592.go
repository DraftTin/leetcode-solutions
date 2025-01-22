package main

import (
	"sort"
)

func maximizeGreatness(nums []int) int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)
	res := 0
	i, j := 0, 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			res++
			i++
			j++
		}
	}
	return res
}
