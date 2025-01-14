package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			var n1, n2, n3 int
			if i-1 >= 0 {
				n1 = dp[i-1][j]
			}
			if j-1 >= 0 {
				n2 = dp[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				n3 = dp[i-1][j-1]
			}
			tmp := 0
			if text1[i] == text2[j] {
				tmp = 1
			}
			dp[i][j] = max(n1, n2, n3+tmp)
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}
