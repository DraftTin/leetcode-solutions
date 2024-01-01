func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true
	for i := n - 2; i >= 0; i-- {
		tmp := false
		for j := 1; i+j < n && j <= nums[i]; j++ {
			tmp = tmp || dp[i+j]
		}
		dp[i] = tmp
	}
	return dp[0]
}
