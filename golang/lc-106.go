package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	nodeVal := postorder[len(postorder)-1]
	node := &TreeNode{Val: nodeVal}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeVal {
			node.Left = buildTree(inorder[:i], postorder[0:i])
			node.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return node
}
