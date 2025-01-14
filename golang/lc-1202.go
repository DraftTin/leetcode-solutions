package main

import "sort"

// Graph UnionFind
func smallestStringWithSwaps(s string, pairs [][]int) string {
	res := []byte{}
	parents := make([]int, len(s))
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}
	for _, pair := range pairs {
		union(parents, pair[0], pair[1])
	}
	unions := make(map[int][]byte)
	for i := 0; i < len(s); i++ {
		p := find(parents, i)
		unions[p] = append(unions[p], s[i])
	}
	for _, val := range unions {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}
	for i := 0; i < len(s); i++ {
		root := find(parents, i)
		res = append(res, unions[root][0])
		unions[root] = unions[root][1:]
	}
	return string(res)
}

func union(parents []int, i, j int) {
	lroot, rroot := find(parents, i), find(parents, j)
	parents[lroot] = parents[rroot]
}

func find(parents []int, i int) int {
	if parents[i] != i {
		parents[i] = find(parents, parents[i])
	}
	return parents[i]
}
