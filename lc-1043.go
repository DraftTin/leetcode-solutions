package main

import "math"

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make(map[int]int)
	return dfs(arr, 0, k, dp)
}

func dfs(arr []int, curPos int, k int, dp map[int]int) int {
	if curPos == len(arr) {
		return 0
	}
	if val, ok := dp[curPos]; ok == true {
		return val
	}
	curMax := math.MinInt32
	ans := math.MinInt32
	for i := curPos; i < curPos+k && i < len(arr); i++ {
		curMax = max(curMax, arr[i])
		val := dfs(arr, i+1, k, dp)
		ans = max(ans, curMax*(i-curPos+1)+val)
	}
	dp[curPos] = ans
	return ans
}
