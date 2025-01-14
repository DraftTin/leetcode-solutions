package main

import "math"

func longestSubarray(nums []int) int {
	res := 0
	maxVal := math.MinInt32
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			k := i + 1
			for k < len(nums) && nums[k] == nums[i] {
				k++
			}
			res = max(res, k-i)
			i = k - 1
		}
	}
	return res
}
