package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	if root == nil {
		return ans
	}
	helper(root, "", &ans)
	return ans
}

func helper(root *TreeNode, curStr string, ans *[]string) {
	curStr = curStr + strconv.Itoa(root.Val)
	if root.Left != nil {
		helper(root.Left, curStr+"->", ans)
	}
	if root.Right != nil {
		helper(root.Right, curStr+"->", ans)
	}
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, curStr)
	}
}
