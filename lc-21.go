/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {  
  p1, p2 := list1, list2
  var head, pre *ListNode
  for p1 != nil && p2 != nil {
    if p1.Val < p2.Val {
      if head == nil {
        head = p1
        pre = p1
      } else {
        pre.Next = p1
        pre = p1
      }
      p1 = p1.Next
    } else {
      if head == nil {
        head = p2
        pre = p2
      } else {
        pre.Next = p2
        pre = p2
      }
      p2 = p2.Next
    }
  }
  var node *ListNode
  if p1 != nil {
    node = p1
  } else {
    node = p2
  }
  if head == nil {
    head = node
  } else {
    pre.Next = node
  }
  return head
}
