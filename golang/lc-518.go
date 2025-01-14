package main

import "sort"

func change(amount int, coins []int) int {
	sort.Ints(coins)
	dp := make([][]int, amount+1)
	visited := make([][]bool, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(coins))
		visited[i] = make([]bool, len(coins))
	}
	return dfs(amount, coins, 0, dp, visited)
}

func dfs(amount int, coins []int, pos int, dp [][]int, visited [][]bool) int {
	if amount == 0 {
		return 1
	}
	if pos == len(coins) {
		return 0
	}
	if visited[amount][pos] == true {
		return dp[amount][pos]
	}
	i := 0
	count := 0
	for i*coins[pos] <= amount {
		count += dfs(amount-i*coins[pos], coins, pos+1, dp, visited)
		i++
	}
	dp[amount][pos] = count
	visited[amount][pos] = true
	return count
}
