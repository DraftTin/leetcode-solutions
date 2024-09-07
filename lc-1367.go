package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	return dfs(head, head, root)
}

func dfs(head, curNode *ListNode, root *TreeNode) bool {
	if curNode == nil {
		return true
	}
	if root == nil {
		return false
	}
	if curNode.Val == root.Val {
		curNode = curNode.Next
	} else if head.Val == root.Val {
		head = head.Next
	} else {
		curNode = head
	}
	return dfs(head, curNode, root.Left) || dfs(head, curNode, root.Right)
}
