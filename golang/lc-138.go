/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
  if head == nil {
    return nil
  }
  mp := map[*Node]int{}
  nodes1 := []*Node{}
  p := head
  for p != nil {
    nodes1 = append(nodes1, p)
    mp[p] = len(nodes1) - 1
    p = p.Next
  }
  nodes2 := make([]*Node, len(nodes1))
  for i := 0; i < len(nodes1); i++ {
    nodes2[i] = &Node{Val: nodes1[i].Val}
  }
  for i := 0; i < len(nodes2); i++ {
    if i < len(nodes1) - 1 {
      nodes2[i].Next = nodes2[i + 1]
    } else {
      nodes2[i].Next = nil
    }
    if nodes1[i].Random == nil {
        nodes2[i].Random = nil
    } else {
        nodes2[i].Random = nodes2[mp[nodes1[i].Random]]
    }
  }   
  return nodes2[0]
}
