package main

func wordBreak(s string, wordDict []string) bool {
	mp := map[string]bool{}
	for _, word := range wordDict {
		mp[word] = true
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if mp[s[:i+1]] == true {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if mp[s[j+1:i+1]] == true && dp[j] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)-1]
}
