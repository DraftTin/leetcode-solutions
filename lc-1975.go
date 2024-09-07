package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)
	minVal := math.MaxInt32
	count := 0
	var sumVal int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := matrix[i][j]
			if matrix[i][j] < 0 {
				tmp = -tmp
				count++
			}
			sumVal += int64(tmp)
			minVal = min(minVal, tmp)
		}
	}
	if count%2 == 0 {
		return sumVal
	}
	return sumVal - 2*int64(minVal)
}
