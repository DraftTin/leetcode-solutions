type Node struct {
	dist        int
	pre         int
	adjacencies []Edge
}

type Edge struct {
	to     int
	weight int
}

func networkDelayTime(times [][]int, n int, k int) int {
	nodes := make([]Node, n)
	for i, _ := range nodes {
		nodes[i].pre = -1
		nodes[i].dist = 6000 * 100
	}
	for _, time := range times {
		nodes[time[0]-1].adjacencies = append(nodes[time[0]-1].adjacencies, Edge{to: time[1] - 1, weight: time[2]})
	}
	nodes[k-1].dist = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		minVal := math.MaxInt32
		index := -1
		for j := 0; j < n; j++ {
			if nodes[j].dist < minVal && visited[j] == false {
				minVal = nodes[j].dist
				index = j
			}
		}
		visited[index] = true
		for _, edge := range nodes[index].adjacencies {
			if nodes[edge.to].dist > nodes[index].dist+edge.weight {
				nodes[edge.to].dist = nodes[index].dist + edge.weight
				nodes[edge.to].pre = index
			}
		}
	}
	ans := 0
	for i, _ := range nodes {
		if nodes[i].pre == -1 && i != k-1 {
			return -1
		}
		if ans < nodes[i].dist {
			ans = nodes[i].dist
		}
	}
	return ans
}


