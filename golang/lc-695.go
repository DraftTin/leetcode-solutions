package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs695(grid, i, j))
			}
		}
	}
	return res
}

func dfs695(grid [][]int, row, col int) int {
	val := 0
	grid[row][col] = 2
	if row < len(grid)-1 && grid[row+1][col] == 1 {
		val += dfs695(grid, row+1, col)
	}

	if col < len(grid[0])-1 && grid[row][col+1] == 1 {
		val += dfs695(grid, row, col+1)
	}
	if row > 0 && grid[row-1][col] == 1 {
		val += dfs695(grid, row-1, col)
	}
	if col > 0 && grid[row][col-1] == 1 {
		val += dfs695(grid, row, col-1)
	}
	return val + 1
}
