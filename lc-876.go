package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	target := n / 2
	p = head
	for i := 0; i < target; i++ {
		p = p.Next
	}
	return p
}
