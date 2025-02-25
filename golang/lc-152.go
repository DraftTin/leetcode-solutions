package main

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxDp1 := make([]int, len(nums))
	maxDp2 := make([]int, len(nums))
	minDp2 := make([]int, len(nums))
	maxDp2[0] = nums[0]
	minDp2[0] = nums[0]
	maxDp1[1] = nums[0]
	maxDp2[1] = max(maxDp2[0]*nums[1], nums[1])
	minDp2[1] = min(minDp2[0]*nums[1], nums[1])
	for i := 2; i < len(nums); i++ {
		maxDp1[i] = max(maxDp1[i-1], maxDp2[i-1])
		maxDp2[i] = max(nums[i], max(nums[i]*maxDp2[i-1], nums[i]*minDp2[i-1]))
		minDp2[i] = min(nums[i], min(nums[i]*maxDp2[i-1], nums[i]*minDp2[i-1]))

	}
	return max(maxDp1[len(nums)-1], maxDp2[len(nums)-1])
}
