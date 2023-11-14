type Node struct {
	cost  int
	edges []Edge
}

type Edge struct {
	destination *Node
	cost        int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Graph struct {
	nodes []*Node
}

func Constructor(n int, edges [][]int) Graph {
	graph := Graph{}
	graph.nodes = make([]*Node, n)
	for i := 0; i < n; i++ {
		graph.nodes[i] = &Node{}
	}
	for _, edge := range edges {
		graph.nodes[edge[0]].edges = append(graph.nodes[edge[0]].edges, Edge{destination: graph.nodes[edge[1]], cost: edge[2]})
	}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	this.nodes[edge[0]].edges = append(this.nodes[edge[0]].edges, Edge{destination: this.nodes[edge[1]], cost: edge[2]})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	pq := make(PriorityQueue, 0)
	visited := make(map[*Node]bool)
	n := len(this.nodes)
	for i := 0; i < n; i++ {
		this.nodes[i].cost = math.MaxInt32
	}
	this.nodes[node1].cost = 0
	heap.Init(&pq)
	heap.Push(&pq, this.nodes[node1])
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		if visited[current] {
			continue
		}
		visited[current] = true
		for _, edge := range current.edges {
			newCost := current.cost + edge.cost
			if newCost < edge.destination.cost {
				edge.destination.cost = newCost
				heap.Push(&pq, edge.destination)
			}
		}
	}
	if this.nodes[node2].cost == math.MaxInt32 {
		return -1
	}
	return this.nodes[node2].cost
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
