package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	helper145(root, &res)
	return res
}

func helper145(node *TreeNode, res *[]int) {
	if node.Left != nil {
		helper145(node.Left, res)
	}
	if node.Right != nil {
		helper145(node.Right, res)
	}
	*res = append(*res, node.Val)
}
