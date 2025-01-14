package main

func maxProfit(prices []int) int {
	n := len(prices)
	dp1, dp2, dp3 := make([]int, n), make([]int, n), make([]int, n)
	dp1[0] = 0
	dp2[0] = -prices[0]
	dp3[0] = 0
	for i := 1; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp3[i-1])
		dp2[i] = max(dp2[i-1], dp1[i-1]-prices[i])
		dp3[i] = dp2[i-1] + prices[i]
	}
	return max(dp1[n-1], dp3[n-1])
}
