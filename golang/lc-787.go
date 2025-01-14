package main

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
		nodes[flight[0]].adjacencies = append(nodes[flight[0]].adjacencies, Edge{to: flight[1], weight: flight[2]})
	}
	queue := [][2]int{[2]int{src, 0}}
	level := 0
	size := 1
	for level <= k && len(queue) != 0 {
		update := queue[0]
		queue = queue[1:]
		if nodes[update[0]].dist > update[1] {
			nodes[update[0]].dist = update[1]
		} else {
			continue
		}
		index := update[0]
		for _, edge := range nodes[index].adjacencies {
			if edge.weight+nodes[index].dist < nodes[edge.to].dist {
				newDist := edge.weight + nodes[index].dist
				queue = append(queue, [2]int{edge.to, newDist})
				if edge.to == dst && ans > newDist {
					ans = newDist
				}
			}
		}
		size--
		if size == 0 {
			size = len(queue)
			level++
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
