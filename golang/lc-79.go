func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, word, 0, i, j, visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, curPos, row, col int, visited [][]bool) bool {
	if curPos == len(word)-1 {
		return true
	}
	visited[row][col] = true
	if row > 0 && board[row-1][col] == word[curPos+1] && !visited[row-1][col] && dfs(board, word, curPos+1, row-1, col, visited) {
		return true
	}
	if col > 0 && board[row][col-1] == word[curPos+1] && !visited[row][col-1] && dfs(board, word, curPos+1, row, col-1, visited) {
		return true
	}
	if row < len(board)-1 && board[row+1][col] == word[curPos+1] && !visited[row+1][col] && dfs(board, word, curPos+1, row+1, col, visited) {
		return true
	}
	if col < len(board[0])-1 && board[row][col+1] == word[curPos+1] && !visited[row][col+1] && dfs(board, word, curPos+1, row, col+1, visited) {
		return true
	}
	visited[row][col] = false
	return false
}
