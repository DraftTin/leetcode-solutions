/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	ans := [][]int{[]int{}}
	size := 1
	for len(queue) != 0 {
		tmp := queue[0]
		ans[len(ans)-1] = append(ans[len(ans)-1], tmp.Val)
		queue = queue[1:]
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		size--
		if size == 0 {
			size = len(queue)
			ans = append(ans, []int{})
		}
	}
	ans = ans[:len(ans)-1]
	return ans
}
