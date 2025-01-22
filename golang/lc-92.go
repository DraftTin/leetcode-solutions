package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	node := &ListNode{}
	node.Next = head
	for i := 1; i < left; i++ {
		node = node.Next
	}
	p1 := node.Next
	var p2, p3 *ListNode
	p2 = p1.next
	if p2 != nil {
		p3 = p2.Next
	}
	for i := left; i < right; i++ {
		p2.Next = p1
		p2 = p3
		if p3 != nil {
			p3 = p3.Next
		}
		p1 = p2
	}
	if node.Next == head {
		head.Next = nil
		return p1
	}
	node.Next.Next = p2
	node.Next = p1
	return head
}
