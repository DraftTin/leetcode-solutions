/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	ans, _ := helper(root, k)
	return ans
}

func helper(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, 0
	}
	lans, lcount := helper(root.Left, k)
	if lcount >= k {
		return lans, lcount
	} else if lcount == k-1 {
		return root.Val, k
	}
	rans, rcount := helper(root.Right, k-lcount-1)
	return rans, rcount + lcount + 1
}
