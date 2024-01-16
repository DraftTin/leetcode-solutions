package main

import "math"

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	dp := make([][]int, n)
	visited := make([][]bool, n)
	visiting := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		visited[i] = make([]bool, m)
		visiting[i] = make([]bool, m)
	}
	dfs(heights, 0, 0, dp, visited, visiting)
	return dp[0][0]
}

func dfs(heights [][]int, row, col int, dp [][]int, visited, visiting [][]bool) {
	if row == len(heights)-1 && col == len(heights[0])-1 {
		return
	}
	if visited[row][col] {
		return
	}
	visiting[row][col] = true
	effort := math.MaxInt32
	if row > 0 && visiting[row-1][col] == false {
		dfs(heights, row-1, col, dp, visited, visiting)
		effort = min(effort, max(dp[row-1][col], abs(heights[row][col], heights[row-1][col])))
	}
	if col > 0 && visiting[row][col-1] == false {
		dfs(heights, row, col-1, dp, visited, visiting)
		effort = min(effort, max(dp[row][col-1], abs(heights[row][col], heights[row][col-1])))
	}
	if row < len(heights)-1 && visiting[row+1][col] == false {
		dfs(heights, row+1, col, dp, visited, visiting)
		effort = min(effort, max(dp[row+1][col], abs(heights[row][col], heights[row+1][col])))
	}
	if col < len(heights[0])-1 && visiting[row][col+1] == false {
		dfs(heights, row, col+1, dp, visited, visiting)
		effort = min(effort, max(dp[row][col+1], abs(heights[row][col], heights[row][col+1])))
	}
	dp[row][col] = effort
	visiting[row][col] = false
	visited[row][col] = true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
