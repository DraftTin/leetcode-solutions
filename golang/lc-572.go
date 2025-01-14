/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	s1, s2 := serialize(root), serialize(subRoot)
	fmt.Println(s1, s2)
	return KMP(s1, s2)
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return "&" + serialize(root.Left) + "," + strconv.Itoa(root.Val) + "," + serialize(root.Right) + "&"

}

func KMP(s1 string, s2 string) bool {
	next := getNext(s2)
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if j == -1 || s1[i] == s2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return j == len(s2)
}

func getNext(s string) []int {
	next := make([]int, len(s))
	i, j := 0, -1
	next[0] = -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
