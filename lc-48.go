package main

// Matrix Manipulation + In-Place
func rotate(matrix [][]int) {
	for i := 0; i < (len(matrix)+1)/2; i++ {
		for j := i; j < len(matrix)-1-i; j++ {
			helper(matrix, i, j, i)
		}
	}
}

func helper(matrix [][]int, startRow, startCol, layer int) {
	nextRow, nextCol := startRow+startCol-layer, len(matrix)-1-layer
	tmp1 := matrix[nextRow][nextCol]
	tmp2 := matrix[startRow][startCol]
	matrix[nextRow][nextCol] = tmp2
	tmp2 = tmp1
	nextRow, nextCol = len(matrix)-1-layer, len(matrix)-1-nextRow
	tmp1 = matrix[nextRow][nextCol]
	matrix[nextRow][nextCol] = tmp2
	tmp2 = tmp1
	nextRow, nextCol = nextCol, layer
	tmp1 = matrix[nextRow][nextCol]
	matrix[nextRow][nextCol] = tmp2
	tmp2 = tmp1
	matrix[startRow][startCol] = tmp2
}
