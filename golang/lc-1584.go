type Edge struct {
	from   int
	to     int
	weight int
}

func union(parents []int, i, j int) {
	p1, p2 := findParent(parents, i), findParent(parents, j)
	if p1 != p2 {
		parents[p2] = p1
	}
}

func findParent(parents []int, i int) int {
	tmp := i
	for parents[i] != i {
		i = parents[i]
	}
	parents[tmp] = i
	return i
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{from: i, to: j, weight: int(math.Abs(float64(points[i][0]-points[j][0]))) + int(math.Abs(float64(points[i][1]-points[j][1])))})
		}
	}
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight <= edges[j].weight
	})
	ans := 0
	for _, edge := range edges {
		p1, p2 := findParent(parents, edge.from), findParent(parents, edge.to)
		if p1 != p2 {
			ans += edge.weight
			union(parents, p1, p2)
		}
	}
	return ans
}
