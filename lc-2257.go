package main

// Matrix Manipulation + Simulation
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for _, wall := range walls {
		row, col := wall[0], wall[1]
		grid[row][col] = 1
	}
	for _, guard := range guards {
		row, col := guard[0], guard[1]
		grid[row][col] = 1
	}
	for _, guard := range guards {
		row, col := guard[0], guard[1]
		for i := row + 1; i < m; i++ {
			if grid[i][col] == 0 {
				grid[i][col] = 2
			} else if grid[i][col] == 1 {
				break
			}
		}
		for i := row - 1; i >= 0; i-- {
			if grid[i][col] == 0 {
				grid[i][col] = 2
			} else if grid[i][col] == 1 {
				break
			}
		}
		for i := col + 1; i < n; i++ {
			if grid[row][i] == 0 {
				grid[row][i] = 2
			} else if grid[row][i] == 1 {
				break
			}
		}
		for i := col - 1; i >= 0; i-- {
			if grid[row][i] == 0 {
				grid[row][i] = 2
			} else if grid[row][i] == 1 {
				break
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
		}
	}
	return res

}
