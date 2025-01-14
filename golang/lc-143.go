/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
  nodes := []*ListNode{}
  for head != nil {
    nodes = append(nodes, head)
    head = head.Next
  }
  i, j := 0, len(nodes) - 1
  var newHead, pre *ListNode
  for i <= j {
    nodes[i].Next = nodes[j]
    if newHead == nil {
      newHead = nodes[i]
    } else {
      pre.Next = nodes[i]
    }
    pre = nodes[j]
    nodes[j].Next = nil
    i++
    j--
  }
  head = newHead
}
