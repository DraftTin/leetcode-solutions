func tictactoe(moves [][]int) string {
	grid := [3][3]int{}
	for index, move := range moves {
		if index%2 == 0 {
			grid[move[0]][move[1]] = 1
		} else {
			grid[move[0]][move[1]] = 2
		}
	}
	res := judge(grid, 0, 0, 0, 1, 0, 2)
	if res != "ok" {
		return res
	}
	res = judge(grid, 1, 0, 1, 1, 1, 2)
	if res != "ok" {
		return res
	}
	res = judge(grid, 2, 0, 2, 1, 2, 2)
	if res != "ok" {
		return res
	}
	res = judge(grid, 0, 0, 1, 0, 2, 0)
	if res != "ok" {
		return res
	}
	res = judge(grid, 0, 1, 1, 1, 2, 1)
	if res != "ok" {
		return res
	}
	res = judge(grid, 0, 2, 1, 2, 2, 2)
	if res != "ok" {
		return res
	}
	res = judge(grid, 0, 0, 1, 1, 2, 2)
	if res != "ok" {
		return res
	}
	res = judge(grid, 2, 0, 1, 1, 0, 2)
	if res != "ok" {
		return res
	}
	if len(moves) != 9 {
		return "Pending"
	}
	return "Draw"
}

func judge(grid [3][3]int, x1, y1, x2, y2, x3, y3 int) string {
	if grid[y1][x1] == grid[y2][x2] && grid[y1][x1] == grid[y3][x3] {
		if grid[y1][x1] == 1 {
			return "A"
		} else if grid[y1][x1] == 2 {
			return "B"
		} else {
			return "Pending"
		}
	}
	return "ok"
}
