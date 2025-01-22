package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	nodeVal := preorder[0]
	node := &TreeNode{Val: nodeVal}
	if len(preorder) < 1 {
		return node
	}
	targetVal := preorder[1]
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == targetVal {
			node.Left = buildTree(preorder[1:i+2], postorder[:i+1])
			node.Right = buildTree(preorder[i+2:], postorder[i+1:len(postorder)-1])
			break
		}
	}
	return node
}
