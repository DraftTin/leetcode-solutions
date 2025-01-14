package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	curNode := head
	size := 0
	for curNode != nil {
		curNode = curNode.Next
		size++
	}
	res := make([]*ListNode, 0, k)
	pieceSize := size / k
	rest := size - pieceSize*k
	curNode = head
	var preNode *ListNode
	for i := 0; i < k; i++ {
		res = append(res, curNode)
		for j := 0; j < pieceSize; j++ {
			preNode = curNode
			curNode = curNode.Next
		}
		if rest > 0 {
			rest--
			preNode = curNode
			curNode = curNode.Next
		}
		if preNode != nil {
			preNode.Next = nil
		}
	}
	return res
}
