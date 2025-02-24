package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper98(root, math.MinInt, math.MaxInt)
}

func helper98(node *TreeNode, leftBound, rightBound int) bool {
	if node == nil {
		return true
	}
	if node.Val <= leftBound || node.Val >= rightBound {
		return false
	}
	return helper98(node.Left, leftBound, node.Val) && helper98(node.Right, node.Val, rightBound)

}
