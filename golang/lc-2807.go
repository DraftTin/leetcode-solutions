package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Euclidean Algorithm Proof:
//	a = k * b + c => num that can divide b and c can also divide a
//	c = a - k * b => num that can divide a and b can also divide c
//	So the gcd of b, c is as the same as that of a, b
//	done.
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		if pre == nil {
			pre = cur
			cur = cur.Next
		} else {
			val := gcd(pre.Val, cur.Val)
			node := &ListNode{Val: val}
			pre.Next = node
			node.Next = cur
			pre = cur
			cur = cur.Next
		}
	}
	return head
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
