package main

import "math"

func minCost(nums []int, costs []int) int64 {
	ascSta, descSta := []int{}, []int{}
	dp := make([]int64, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for len(ascSta) != 0 && nums[i] >= nums[ascSta[len(ascSta)-1]] {
			tmp := ascSta[len(ascSta)-1]
			ascSta = ascSta[:len(ascSta)-1]
			dp[i] = min(dp[i], dp[tmp]+int64(costs[i]))
		}
		for len(descSta) != 0 && nums[i] < nums[descSta[len(descSta)-1]] {
			tmp := descSta[len(descSta)-1]
			descSta = descSta[:len(descSta)-1]
			dp[i] = min(dp[i], dp[tmp]+int64(costs[i]))
		}
		ascSta = append(ascSta, i)
		descSta = append(descSta, i)
	}
	return dp[len(nums)-1]
}
