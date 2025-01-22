package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	helper144(root, res)
	return res
}

func helper144(node *TreeNode, res []int) {
	res = append(res, node.Val)
	if node.Left != nil {
		helper144(node, res)
	}
	if node.Right != nil {
		helper144(node, res)
	}
}
