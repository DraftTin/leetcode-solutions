package main

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	if text1[0] == text2[0] {
		dp[0][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i > 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j])
				}
				if j > 0 {
					dp[i][j] = max(dp[i][j], dp[i][j-1])
				}
			}
		}
	}
	return dp[n-1][m-1]
}
