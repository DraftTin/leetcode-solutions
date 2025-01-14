package main

type UnionNode struct {
	Parent *UnionNode
	Size   int
	Rank   int
	Row    int
	Col    int
}

func removeStones(stones [][]int) int {
	colMap, rowMap := map[int][]*UnionNode{}, map[int][]*UnionNode{}
	nodeList := []*UnionNode{}
	for _, stone := range stones {
		node := &UnionNode{Size: 1, Rank: 1, Row: stone[0], Col: stone[1]}
		nodeList = append(nodeList, node)
		node.Parent = node
		rowMap[stone[0]] = append(rowMap[stone[0]], node)
		colMap[stone[1]] = append(colMap[stone[1]], node)
	}
	for _, node := range nodeList {
		for _, val := range rowMap[node.Row] {
			if val != node {
				p1, p2 := find947(val), find947(node)
				if p1 != p2 {
					union947(p1, p2)
				}
				break
			}
		}
		for _, val := range colMap[node.Col] {
			if val != node {
				p1, p2 := find947(val), find947(node)
				if p1 != p2 {
					union947(p1, p2)
				}
				break
			}
		}
	}
	groupMap := map[*UnionNode]int{}
	for _, node := range nodeList {
		parent := find947(node)
		groupMap[parent] = parent.Size
	}
	res := 0
	for _, val := range groupMap {
		res += val - 1
	}
	return res
}

func union947(node1, node2 *UnionNode) {
	p1, p2 := find947(node1), find947(node2)
	if p1.Rank < p2.Rank {
		p1.Parent = p2
		p2.Size += p1.Size
	} else {
		p2.Parent = p1
		p1.Size += p2.Size
		if p1.Rank == p2.Rank {
			p2.Rank++
		}
	}

}

func find947(node *UnionNode) *UnionNode {
	if node.Parent != node {
		node.Parent = find947(node.Parent)
	}
	return node.Parent
}
