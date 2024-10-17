package main

import "fmt"

func setZeroes(matrix [][]int) {
	mark := make(map[string]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 && mark[fmt.Sprintf("%s,%s", i, j)] == false {
				if mark[fmt.Sprintf("%sr", i)] == false {
					for k := 0; k < len(matrix[0]); k++ {
						matrix[i][k] = 0
						mark[fmt.Sprintf("%s,%s", i, k)] = true
					}
				}
				if mark[fmt.Sprintf("%sc", i)] == false {
					for k := 0; k < len(matrix); k++ {
						matrix[k][j] = 0
						mark[fmt.Sprintf("%s,%s", k, j)] = true
					}
				}
			}
		}
	}
}
