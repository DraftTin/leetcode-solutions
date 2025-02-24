package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	nodeList := []*ListNode{}
	p := head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}
	i, j := 0, len(nodeList)-1
	var pre *ListNode
	for i <= j {
		nodeList[i].Next = nodeList[j]
		if pre != nil {
			pre.Next = nodeList[i]
		}
		pre = nodeList[j]
		nodeList[j].Next = nil
		i++
		j--
	}
	head = nodeList[0]
}
