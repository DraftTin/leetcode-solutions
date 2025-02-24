package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	que := []*TreeNode{root}
	size := 1
	res := []int{}
	for len(que) != 0 {
		tmp := que[0]
		que = que[1:]
		if tmp.Left != nil {
			que = append(que, tmp.Left)
		}
		if tmp.Right != nil {
			que = append(que, tmp.Right)
		}
		size--
		if size == 0 {
			res = append(res, tmp.Val)
			size = len(que)
		}
	}
	return res
}
