/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, ans := maxDepth(root)
	return ans
}

func maxDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftAns := maxDepth(root.Left)
	rightDepth, rightAns := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1, leftAns && rightAns && leftDepth-rightDepth >= -1 && leftDepth-rightDepth <= 1
}
