package main

// DP
func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 0; i < target+1; i++ {
		dp[0][i] = false
	}
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		for j := 1; j < target+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				if dp[i-1][j-nums[i-1]] == true {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[len(dp)-1][target]
}
