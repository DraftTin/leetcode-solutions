func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		minSteps := math.MaxInt32
		for j := 1; i+j < n && j <= nums[i]; j++ {
			minSteps = min(minSteps, dp[i+j])
		}
		if minSteps == math.MaxInt32 {
			dp[i] = math.MaxInt32
		} else {
			dp[i] = minSteps + 1
		}
	}
	return dp[0]
}

