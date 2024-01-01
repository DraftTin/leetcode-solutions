// if we use BFS and Dijkstra to solve a shortest-path problem, we shouldn't reflect the update of the next
// level into this level
type Node struct {
	dist        int
	adjacencies []Edge
}

type Edge struct {
	to     int
	weight int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	ans := math.MaxInt32
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].dist = math.MaxInt32
	}
	for _, flight := range flights {
		nodes[flight[0]-1].adjacencies = append(nodes[i].adjacencies, Edge{to: flight[1] - 1, weight: flight[2]})
	}
	queue := []int{k - 1}
	level := 0
	size := 1
	for level <= k && len(queue) != 0 {
		index := queue[0]
		queue = queue[1:]
		if index == k-1 && ans > nodes[index].dist {
			ans = nodes[index].dist
		}
		for _, edge := range nodes[index].adjacencies {
			if edge.weight+nodes[index].dist < nodes[edge.to].dist {
				nodes[edge.to].dist = edge.weight + nodes[index].dist
				queue = append(queue, edge.to)
			}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
