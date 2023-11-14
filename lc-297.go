/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	vals := preorderTraversal(root)
	return strings.Join(vals, " ")
}

func preorderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{"#"}
	}
	vals := []string{strconv.Itoa(root.Val)}
	leftArr := preorderTraversal(root.Left)
	rightArr := preorderTraversal(root.Right)
	vals = append(vals, leftArr...)
	vals = append(vals, rightArr...)
	return vals
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	index := 0
	vals := strings.Split(data, " ")
	return helper(vals, &index)
}

func helper(vals []string, index *int) *TreeNode {
	if vals[*index] == "#" {
		*index += 1
		return nil
	}
	val, _ := strconv.Atoi(vals[*index])
	node := &TreeNode{Val: val}
	*index += 1
	node.Left = helper(vals, index)
	node.Right = helper(vals, index)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
