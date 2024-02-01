package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		if i+2 < len(nums) {
			if nums[i+2]-nums[i] > k {
				return [][]int{}
			} else {
				ans = append(ans, nums[i:i+3])
			}
		} else if i+2 >= len(nums) {
			if nums[len(nums)-1]-nums[i] > k {
				return [][]int{}
			} else {
				ans = append(ans, nums[i:len(nums)])
			}
		}
	}
	return ans
}
