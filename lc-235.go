package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	f1, f2 := -1, -1
	if root == p {
		f1 = 0
	} else if findNode(root.Left, p) == true {
		f1 = 1
	} else {
		f1 = 2
	}
	if root == q {
		f2 = 0
	} else if findNode(root.Left, q) == true {
		f2 = 1
	} else {
		f2 = 2
	}
	if f1 != f2 {
		ans = root
	}
	return ans
}

func findNode(root *TreeNode, node *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	return findNode(root.Left, node) || findNode(root.Right, node)
}
