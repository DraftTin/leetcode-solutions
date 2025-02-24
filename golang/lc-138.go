package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	p := head
	for p != nil {
		node := &Node{Val: p.Val}
		mp[p] = node
		p = p.Next
	}
	newHead := mp[head]
	p = head
	newP := newHead
	for p != nil {
		newP.Next = mp[p.Next]
		newP.Random = mp[p.Random]
		newP = newP.Next
		p = p.Next
	}
	return newHead
}
