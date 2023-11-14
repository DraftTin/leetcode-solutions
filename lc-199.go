/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	size := 1
	queue := []*TreeNode{root}
	ans := []int{}
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		size--
		if size == 0 {
			ans = append(ans, tmp.Val)
			size = len(queue)
		}
	}
	return ans
}
