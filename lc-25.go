/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
  pre := head
  size := 0
  for pre != nil {
    pre = pre.Next
    size++
  }
  var newHead *ListNode
  for i := 0; i < size / k; i++ {
    var p1, p2, p3 *ListNode 
    p1 = head
    if p1 != nil {
      p2 = p1.Next
    }
    if p2 != nil {
      p3 = p2.Next
    }
    for j := 0; j < k - 1; j++ {
      p2.Next, p1, p2 = p1, p2, p3
      if p3 != nil {
        p3 = p3.Next
      }
    }
    if newHead == nil {
      newHead = p1
    } else {
      pre.Next = p1
    }
    pre = head
    head = p2
    pre.Next = head
  }
  return newHead
}
