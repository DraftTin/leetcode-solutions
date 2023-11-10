/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
  if len(lists) == 0 {
    return nil
  }
  head := lists[0]
  for i := 1; i < len(lists); i++ {
    head = mergeTwoLists(head, lists[i])
  }
  return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  var head *ListNode
  var pre *ListNode
  for list1 != nil && list2 != nil {
    var node *ListNode
    if list1.Val < list2.Val {
      node = list1
      list1 = list1.Next
    } else {
      node = list2
      list2 = list2.Next
    }
    if head == nil {
      head = node
      pre = node
    } else {
      pre.Next = node
      pre = node
    }
  }
  node := list1
  if list2 != nil {
    node = list2
  }
  if head == nil {
    head = node
    return head  
  }
  pre.Next = node
  return head
}
