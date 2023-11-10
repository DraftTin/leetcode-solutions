/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root.Right != nil {
		invertTree(root.Right)
	}
	if root.Left != nil {
		inverTree(root.Left)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}   
