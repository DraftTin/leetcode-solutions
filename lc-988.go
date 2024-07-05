package main

func update(cur, res []int) []int {
	if len(res) == 0 {
		return append([]int{}, cur...)
	}

	l, r := len(cur)-1, len(res)-1

	for l >= 0 && r >= 0 {
		if cur[l] < res[r] {
			return append([]int{}, cur...)
		} else if cur[l] > res[r] {
			return res
		}

		l--
		r--
	}

	if r >= 0 {
		return append([]int{}, cur...)
	}
	return res
}

func dfs(node *TreeNode, cur []int, res []int) []int {
	if node == nil {
		return res
	}

	cur = append(cur, node.Val)

	if node.Left == nil && node.Right == nil {
		res = update(cur, res)
	} else {
		res = dfs(node.Left, cur, res)
		res = dfs(node.Right, cur, res)
	}

	cur = cur[:len(cur)-1]
	return res
}

func smallestFromLeaf(root *TreeNode) string {
	res := []int{}
	cur := []int{}

	res = dfs(root, cur, res)

	str := []byte{}
	for i := len(res) - 1; i >= 0; i-- {
		str = append(str, byte('a'+res[i]))
	}

	return string(str)
}
