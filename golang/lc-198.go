package main

func rob(nums []int) int {
	dp1, dp2 := make([]int, len(nums)), make([]int, len(nums))
	dp1[0] = 0
	dp2[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp1[i] = dp2[i-1]
		dp2[i] = dp1[i-1] + nums[i]
	}
	return max(dp1[len(dp1)-1], dp2[len(dp2)-1])
}
