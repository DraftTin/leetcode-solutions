/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }
  var p1, p2, p3 *ListNode
  p1 = head
  p2 = p1.Next
  if p2 != nil {
    p3 = p2.Next
  }
  for p2 != nil {
    p2.Next = p1
    p1 = p2
    p2 = p3
    if p3 != nil {
      p3 = p3.Next
    }
  }
  head.Next = nil
  return p1
}
