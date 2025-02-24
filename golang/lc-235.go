package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper235(root, p, q)
}

func helper235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	l := helper235(root.Left, p, q)
	r := helper235(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l == nil {
		return r
	}
	return l
}
