/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	findTwoNodes(root, p, q, &ans)
	return ans
}

func findTwoNodes(root, p, q *TreeNode, ans **TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	}
	found1, found2 := false, false
	if root == p {
		found1 = true
	}
	if root == q {
		found2 = true
	}
	f1, f2 := findTwoNodes(root.Left, p, q, ans)
	if *ans != nil {
		return true, true
	}
	found1 = found1 || f1
	found2 = found2 || f2
	f1, f2 = findTwoNodes(root.Right, p, q, ans)
	if *ans != nil {
		return true, true
	}
	found1 = found1 || f1
	found2 = found2 || f2
	if found1 && found2 {
		*ans = root
	}
	return found1, found2
}
