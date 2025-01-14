/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
  if head == nil {
    return false
  }
  fast, slow := head, head.Next
  for fast != nil && fast != slow {
    fast = fast.Next
    slow = slow.Next
    if fast != nil {
      fast = fast.Next
    }
  }
  return fast == nil
}
