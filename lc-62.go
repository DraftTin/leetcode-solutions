package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1
	dfs(dp, 0, 0)
	return dp[0][0]
}

func dfs(dp [][]int, row, col int) {
	if dp[row][col] != 0 {
		return
	}
	count := 0
	if row < len(dp)-1 {
		dfs(dp, row+1, col)
		count += dp[row+1][col]
	}
	if col < len(dp[0])-1 {
		dfs(dp, row, col+1)
		count += dp[row][col+1]
	}
	dp[row][col] = count
}
