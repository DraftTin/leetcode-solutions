package main

func nearestExit(maze [][]byte, entrance []int) int {
	que := [][]int{}
	row := entrance[0]
	col := entrance[1]
	visited := map[[2]int]bool{}
	visited[[2]int{row, col}] = true
	if row-1 >= 0 && maze[row-1][col] == '.' {
		que = append(que, []int{row - 1, col})
	}
	if row+1 < len(maze) && maze[row+1][col] == '.' {
		que = append(que, []int{row + 1, col})
	}
	if col-1 >= 0 && maze[row][col-1] == '.' {
		que = append(que, []int{row, col - 1})
	}
	if col+1 < len(maze[0]) && maze[row][col+1] == '.' {
		que = append(que, []int{row, col + 1})
	}
	size := len(que)
	level := 1
	for len(que) > 0 {
		row, col := que[0][0], que[0][1]
		que = que[1:]
		if row == 0 || col == 0 || row == len(maze)-1 || col == len(maze[0])-1 {
			return level
		}
		if row-1 >= 0 && maze[row-1][col] == '.' && !visited[[2]int{row - 1, col}] {
			visited[[2]int{row - 1, col}] = true
			que = append(que, []int{row - 1, col})
		}
		if row+1 < len(maze) && maze[row+1][col] == '.' && !visited[[2]int{row + 1, col}] {
			visited[[2]int{row + 1, col}] = true
			que = append(que, []int{row + 1, col})
		}
		if col-1 >= 0 && maze[row][col-1] == '.' && !visited[[2]int{row, col - 1}] {
			visited[[2]int{row, col - 1}] = true
			que = append(que, []int{row, col - 1})
		}
		if col+1 >= 0 && maze[row][col+1] == '.' && !visited[[2]int{row, col + 1}] {
			visited[[2]int{row, col + 1}] = true
			que = append(que, []int{row, col + 1})
		}
		size--
		if size == 0 {
			size = len(que)
			level++
		}
	}
	return -1

}
