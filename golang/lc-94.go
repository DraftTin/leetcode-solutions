package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	helper94(root, &res)
	return res
}

func helper94(node *TreeNode, res *[]int) {
	if node.Left != nil {
		helper94(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		helper94(node.Right, res)
	}
}
