package main

func maxProfit(prices []int) int {
	dp1 := make([]int, len(prices))
	dp2 := make([]int, len(prices))
	dp3 := make([]int, len(prices))
	dp1[0] = 0
	dp2[0] = -prices[0]
	dp3[0] = 0
	for i := 1; i < len(prices); i++ {
		dp1[i] = max(dp1[i-1], dp3[i-1])
		dp2[i] = max(dp2[i-1], dp1[i-1]-prices[i])
		dp3[i] = dp2[i-1] + prices[i]
	}
	return max(dp1[len(prices)-1], dp3[len(prices)-1])
}
