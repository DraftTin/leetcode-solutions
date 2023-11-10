/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	maxDepth(root, &ans)
	return ans
}

func maxDepth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	rightDepth := maxDepth(root.Right, ans)
	leftDepth := maxDepth(root.Left, ans)
	if rightDepth+leftDepth > *ans {
		*ans = rightDepth + leftDepth
	}
	return max(rightDepth, leftDepth) + 1
}
