package main

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if s[0] != '0' {
		dp[0] = 1
	} else {
		return 0
	}
	if len(s) >= 2 {
		if s[1] != '0' {
			dp[1] = 1
		}
		if canMatch(s[0], s[1]) {
			dp[1]++
		}
	}
	for i := 2; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}
		if canMatch(s[i-1], s[i]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)-1]
}

func canMatch(a, b byte) bool {
	if a != '1' && a != '2' {
		return false
	}
	if a == '1' {
		return true
	}
	return b >= '0' && b <= '6'
}
