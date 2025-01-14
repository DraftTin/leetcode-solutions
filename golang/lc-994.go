func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				visited[i][j] = true
				q = append(q, [2]int{i, j})
			}
		}
	}
	size := len(q)
	ans := 0
	for len(q) != 0 {
		row, col := q[0][0], q[0][1]
		q = q[1:]
		if row > 0 && grid[row-1][col] == 1 && !visited[row-1][col] {
			visited[row-1][col] = true
			q = append(q, [2]int{row - 1, col})
		}
		if col > 0 && grid[row][col-1] == 1 && !visited[row][col-1] {
			visited[row][col-1] = true
			q = append(q, [2]int{row, col - 1})
		}
		if row < len(grid)-1 && grid[row+1][col] == 1 && !visited[row+1][col] {
			visited[row+1][col] = true
			q = append(q, [2]int{row + 1, col})
		}
		if col < len(grid[0])-1 && grid[row][col+1] == 1 && !visited[row][col+1] {
			visited[row][col+1] = true
			q = append(q, [2]int{row, col + 1})
		}
		size--
		if size == 0 {
			size = len(q)
			ans++
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] == false && grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans - 1
}
