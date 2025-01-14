/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var l3 *ListNode
  var p *ListNode
  complement := 0
  for l1 != nil && l2 != nil {
    ans := (l1.Val + l2.Val + complement) % 10 
    if l3 == nil {
      l3 = &ListNode{Val: ans}
      p = l3
    } else {
      p.Next = &ListNode{Val: ans}
      p = p.Next
    }
    complement = (l1.Val + l2.Val + complement) / 10
    l1 = l1.Next
    l2 = l2.Next
  }
  node := l2
  if l1 != nil {
    node = l1
  }
  for node != nil {
    ans := (node.Val + complement) % 10
    complement = (node.Val + complement) / 10
    p.Next = &ListNode{Val: ans}
    p = p.Next
    node = node.Next
  }
  if complement == 1 {
    p.Next = &ListNode{Val: 1}
  }
  return l3
}
