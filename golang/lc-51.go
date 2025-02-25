package main

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	curBoard := make([][]byte, n)
	for i := range curBoard {
		curBoard[i] = make([]byte, n)
		for j := range curBoard[i] {
			curBoard[i][j] = '.'
		}
	}
	dfs51(curBoard, 0, map[int]bool{}, map[int]bool{}, map[int]bool{}, &ans)
	return ans
}

func dfs51(curBoard [][]byte, row int, colMp, slashMap, backSlashMap map[int]bool, ans *[][]string) {
	if row == len(curBoard) {
		res := make([]string, 0, len(curBoard))
		for _, row := range curBoard {
			res = append(res, string(row))
		}
		*ans = append(*ans, res)
	}
	for i := 0; i < len(curBoard); i++ {
		if colMp[i] || slashMap[row-i] || backSlashMap[row+i] {
			continue
		}
		curBoard[row][i] = 'Q'
		colMp[i] = true
		slashMap[row-i] = true
		backSlashMap[row+i] = true
		dfs51(curBoard, row+1, colMp, slashMap, backSlashMap, ans)
		curBoard[row][i] = '.'
		colMp[i] = false
		slashMap[row-i] = false
		backSlashMap[row+i] = false
	}
}
