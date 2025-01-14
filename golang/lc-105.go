/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, 0, inorder, 0, len(inorder)-1)
}

func helper(preorder []int, start int, inorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	node := &TreeNode{Val: preorder[start]}
	index := -1
	for i := left; i < len(inorder); i++ {
		if inorder[i] == preorder[start] {
			index = i
			break
		}
	}
	node.Left = helper(preorder, start+1, inorder, left, index-1)
	node.Right = helper(preorder, start+1+index-left, inorder, index+1, right)
	return node
}
