package main

func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges))
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}
	index := -1
	for i, edge := range edges {
		lroot, rroot := findRoot(parents, edge[0]-1), findRoot(parents, edge[1]-1)
		if lroot == rroot {
			index = i
		}
		union(parents, lroot, rroot)
	}
	return edges[index]
}

func union(parents []int, i, j int) {
	lroot, rroot := findRoot(parents, i), findRoot(parents, j)
	parents[lroot] = rroot
}

func findRoot(parents []int, i int) int {
	if parents[i] != i {
		parents[i] = findRoot(parents, parents[i])
	}
	return parents[i]
}
