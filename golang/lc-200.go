package main

func numIslands(grid [][]byte) int {
	count := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				count++
				dfs(grid, i, j, visited)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, row, col int, visited [][]bool) {
	visited[row][col] = true
	if row > 0 && !visited[row-1][col] && grid[row-1][col] == '1' {
		dfs(grid, row-1, col, visited)
	}
	if col > 0 && !visited[row][col-1] && grid[row][col-1] == '1' {
		dfs(grid, row, col-1, visited)
	}
	if col < len(grid[0])-1 && !visited[row][col+1] && grid[row][col+1] == '1' {
		dfs(grid, row, col+1, visited)
	}
	if row < len(grid)-1 && !visited[row+1][col] && grid[row+1][col] == '1' {
		dfs(grid, row+1, col, visited)
	}
}
