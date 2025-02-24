package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	_, ans := helper(root, k)
	return ans
}

func helper(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, 0
	}
	lnum, lval := helper(root.Left, k)
	if lnum+1 == k {
		return lnum + 1, root.Val
	} else if lnum >= k {
		return lnum + 1, lval
	}
	rnum, rval := helper(root.Right, k-lnum-1)
	return rnum + lnum + 1, rval
}
