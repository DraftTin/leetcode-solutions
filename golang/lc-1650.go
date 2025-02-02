package main

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

// Binary Tree
func lowestCommonAncestor(p *Node, q *Node) *Node {
	root := findRoot(p)
	_, res := search1650(root, p, q)
	return res
}

func findRoot(p *Node) *Node {
	if p.Parent == nil {
		return p
	}
	return findRoot(p.Parent)
}

func search1650(node *Node, p *Node, q *Node) (int, *Node) {
	if node == nil {
		return 0, nil
	}
	res := 0
	if node.Val == p.Val {
		res++
	}
	if node.Val == q.Val {
		res++
	}
	lnum, lnode := search1650(node.Left, p, q)
	rnum, rnode := search1650(node.Right, p, q)
	if lnum == 2 {
		return 2, lnode
	}
	if rnum == 2 {
		return 2, rnode
	}
	return lnum + rnum + res, node
}
