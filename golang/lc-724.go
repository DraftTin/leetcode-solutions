package main

func pivotIndex(nums []int) int {
	sumVal := 0
	leftSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumVal += nums[i]
		leftSum[i] = sumVal
	}
	if sumVal-nums[0] == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if leftSum[i-1] == sumVal-nums[i]-leftSum[i-1] {
			return i
		}
	}
	return -1
}
