package main

func findMaxAverage(nums []int, k int) float64 {
	curSum := 0
	maxSum := 0
	for i := 0; i < k; i++ {
		curSum += nums[i]
	}
	maxSum = curSum
	for i := k; i < len(nums); i++ {
		curSum += nums[i] - nums[i-k]
		maxSum = max(maxSum, curSum)
	}
	return float64(maxSum) / float64(k)
}
