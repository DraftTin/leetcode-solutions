package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	size := 0
	for p != nil {
		size++
		p = p.Next
	}
	if n == size {
		return head.Next
	}
	p = head
	var pre *ListNode
	i := 1
	for p != nil {
		if i == size-n+1 {
			pre.Next = p.Next
			return head
		}
		pre = p
		p = p.Next
		i++
	}
	return head
}
