package main

import "math"

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(dp)-1] = 0
	for i := len(dp) - 2; i >= 0; i-- {
		minVal := math.MaxInt32
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			minVal = min(minVal, dp[i+j])
		}
		if minVal == math.MaxInt32 {
			dp[i] = math.MaxInt32
		} else {
			dp[i] = minVal + 1
		}
	}
	return dp[0]
}
