package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	sta := []*ListNode{}
	p := head
	for p != nil {
		if p.Val == 0 {
			p = p.Next
			continue
		}
		if len(sta) == 0 {
			sta = append(sta, p)
			p = p.Next
			continue
		}
		tmp := p.Val
		i := len(sta) - 1
		for i >= 0 {
			tmp += sta[i].Val
			if tmp == 0 {
				sta = sta[:i]
				break
			}
			i--
		}
		if tmp != 0 {
			sta = append(sta, p)
		}
		p = p.Next
	}
	if len(sta) != 0 {
		var pre *ListNode
		for _, node := range sta {
			if pre == nil {
				pre = node
			} else {
				pre.Next = node
				pre = node
			}
		}
		pre.Next = nil
	} else {
		return nil
	}
	return sta[0]
}
