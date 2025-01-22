package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dict := make(map[*ListNode]bool)
	node := headA
	for node != nil {
		dict[node] = true
	}
	node = headB
	for node != nil {
		if dict[node] == true {
			return node
		}
	}
	return nil
}
