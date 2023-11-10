/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
  if head == nil {
    return nil
  } 
  slow, fast := head, head
  for {
    fast = fast.Next
    if fast != nil {
      fast = fast.Next
    }
    slow = slow.Next
    if fast == nil || fast == slow {
      break
    }
  }
  if fast == nil {
    return nil
  }
  fast = head
  for fast != slow {
    fast = fast.Next
    slow = slow.Next
  }
  return fast
}
