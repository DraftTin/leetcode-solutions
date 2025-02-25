package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	res := 0
	dp1, dp2 := make([]int, len(nums)), make([]int, len(nums))
	dp1[1], dp2[1] = nums[0], nums[0]
	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-1], dp2[i-1])
		dp2[i] = max(dp2[i-1], dp1[i-1]+nums[i])
	}
	res = max(dp1[len(nums)-2], dp2[len(nums)-2])
	dp1, dp2 = make([]int, len(nums)), make([]int, len(nums))
	dp1[1] = 0
	dp2[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp1[i] = max(dp1[i-1], dp2[i-1])
		dp2[i] = max(dp2[i-1], dp1[i-1]+nums[i])
	}
	res = max(res, max(dp1[len(nums)-1], dp2[len(nums)-1]))
	return res
}
