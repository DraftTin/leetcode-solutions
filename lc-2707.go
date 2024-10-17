package main

import "math"

// DP problem
// dp[ind] represents that the result value when start index == ind
// dp[ind] = min(min(dp[ind + k]), 1 + dp[ind + 1]), when s[ind : ind + k + 1] is in dictionary
func minExtraChar(s string, dictionary []string) int {
	dictSet := map[string]bool{}
	for _, str := range dictionary {
		dictSet[str] = true
	}
	dp := make([]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	return dpForMinExtraChar(s, 0, dp, dictSet)
}

func dpForMinExtraChar(s string, ind int, dp []int, dictSet map[string]bool) int {
	if ind >= len(s) {
		return 0
	}
	if dp[ind] != -1 {
		return dp[ind]
	}
	tmp := math.MaxInt32
	for i := ind; i < len(s); i++ {
		if dictSet[s[ind:i+1]] {
			tmp = min(tmp, dpForMinExtraChar(s, i+1, dp, dictSet))
		}
	}
	dp[ind] = tmp
	dp[ind] = min(dp[ind], 1+dpForMinExtraChar(s, ind+1, dp, dictSet))
	return dp[ind]
}
