package main

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			tmp := j
			dp[i][j] = dp[i-1][tmp]
			for tmp >= coins[i-1] {
				tmp = tmp - coins[i-1]
				dp[i][j] += dp[i-1][tmp]
			}
		}
	}
	return dp[len(coins)][amount]
}
