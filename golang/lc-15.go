package main

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	set := map[string]bool{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		target := -num1
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				str := strconv.Itoa(num1) + "," + strconv.Itoa(nums[l]) + "," + strconv.Itoa(nums[r])
				if set[str] == false {
					res = append(res, []int{num1, nums[l], nums[r]})
					set[str] = true
				}
				l++
				r--
			}
		}
	}
	return res
}
