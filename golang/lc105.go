/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	nodeVal := preorder[0]
	node := &TreeNode{Val: nodeVal}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeVal {
			node.Left = buildTree(preorder[1:i+1], inorder[:i])
			node.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return node
}

