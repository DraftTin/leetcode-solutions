package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Linked List Manipulation
func modifiedList(nums []int, head *ListNode) *ListNode {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	fakeHead := &ListNode{}
	fakeHead.Next = head
	pre, cur := fakeHead, head
	for cur != nil {
		if numMap[cur.Val] {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return fakeHead.Next
}
