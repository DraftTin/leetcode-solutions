package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Value-Changed Version
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		node := minimum(root.Right)
		root.Val = node.Val
		root.Right = deleteNode(root.Right, node.Val)
	}
	return root
}

func minimum(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return minimum(root.Left)
}
