package main

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr))
	base := 0
	for i := 0; i < len(arr); i++ {
		base = base ^ arr[i]
		xors[i] = base
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		tmp := xors[query[1]]
		if query[0] != 0 {
			tmp = tmp ^ xors[query[0]-1]
		}
		res[i] = tmp
	}
	return res
}
