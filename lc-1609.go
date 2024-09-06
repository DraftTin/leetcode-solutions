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

// BFS
func isEvenOddTree(root *TreeNode) bool {
	que := []*TreeNode{}
	que = append(que, root)
	size := 1
	level := 0
	prev := math.MinInt32
	for len(que) != 0 {
		node := que[0]
		if level%2 == 0 {
			if node.Val <= prev || node.Val%2 == 0 {
				return false
			}
		}
		if level%2 == 1 {
			if node.Val >= prev || node.Val%2 == 1 {
				return false
			}
		}
		prev = node.Val
		que = que[1:]
		size--
		if node.Left != nil {
			que = append(que, node.Left)
		}
		if node.Right != nil {
			que = append(que, node.Right)
		}
		if size == 0 {
			size = len(que)
			if level%2 == 0 {
				prev = math.MaxInt32
			} else {
				prev = math.MinInt32
			}
			level++
		}
	}
	return true
}
