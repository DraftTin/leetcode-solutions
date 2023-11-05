/**
 * definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
  ans := 0
  helper(nil, nil, root, &ans)
  return ans
}

func helper(grandparent *TreeNode, parent *TreeNode, node *TreeNode, ans *int) {
  if grandparent != nil && grandparent.Val % 2 == 0 {
    (*ans) += node.Val
  }
  if node.Left != nil {  
    helper(parent, node, node.Left, ans)
  }
  if node.Right != nil {  
    helper(parent, node, node.Right, ans)
  }
}
