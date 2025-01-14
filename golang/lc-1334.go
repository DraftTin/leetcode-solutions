package main

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				dist[i][j] = 1000001
			}
		}
	}
	for _, edge := range edges {
		dist[edge[0]][edge[1]] = edge[2]
		dist[edge[1]][edge[0]] = edge[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	minVal := math.MaxInt32
	res := 0
	for i := 0; i < n; i++ {
		tmp := 0
		for j := 0; j < n; j++ {
			if dist[i][j] <= distanceThreshold && i != j {
				tmp++
			}
		}
		if tmp <= minVal {
			minVal = tmp
			res = i
		}
	}
	return res
}
