/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[int]*Node{}
	newNode := &Node{Val: node.Val}
	dfs(node, newNode, visited)
	return newNode
}

func dfs(node1 *Node, node2 *Node, visited map[int]*Node) {
	neighbors := []*Node{}
	visited[node1.Val] = node2
	for _, neighbor := range node1.Neighbors {
		if node, ok := visited[neighbor.Val]; ok == true {
			neighbors = append(neighbors, node)
		} else {
			neighbors = append(neighbors, &Node{Val: neighbor.Val})
			dfs(neighbor, neighbors[len(neighbors)-1], visited)
		}
	}
	node2.Neighbors = neighbors
}
