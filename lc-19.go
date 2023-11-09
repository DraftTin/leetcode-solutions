/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  size := 0
  p := head
  for p != nil {
    size++
    p = p.Next
  }
  var pre *ListNode
  p = head
  k := 1
  n = size - n + 1
  for p != nil {
    if k == n {
      if pre == nil {
        return p.Next
      }
      pre.Next = p.Next
      break
    }
    pre = p
    p = p.Next
    k++
  }
  return head
}
