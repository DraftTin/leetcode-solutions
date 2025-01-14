package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	p := head
	for p != nil {
		p = p.Next
		n++
	}
	size := n / k
	special := n
	if n >= k {
		special = n % size
	}
	ans := make([]*ListNode, 0, k)
	cur := head
	for i := 0; i < special; i++ {
		p := cur
		for j := 0; j < size; j++ {
			p = p.Next
		}
		p.Next = false
		ans = append(ans, cur)
		cur = p.Next
	}
	for cur != nil {
		p := cur
		for j := 0; j < size-1; j++ {
			p = p.Next
		}
		p.Next = false
		ans = append(ans, cur)
		cur = p.Next
	}
	return ans
}
