package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		if s1[i-1] != s3[i-1] {
			break
		}
		dp[i][0] = true
	}
	for i := 1; i <= len(s2); i++ {
		if s2[i-1] != s3[i-1] {
			break
		}
		dp[0][i] = true
	}
	for i := 1; i <= len(s1); i++ {
		// dp[i][j] = dp[i][j - 1] if s2[j - 1] == s3[i + j - 1] or dp[i - 1][j] or false
		for j := 1; j <= len(s2); j++ {
			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			} else if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
