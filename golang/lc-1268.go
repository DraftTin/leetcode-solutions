package main

type TrieNode struct {
	EndNode  bool
	Children [26]*TrieNode
	Val      string
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := &TrieNode{}
	root.Children = [26]*TrieNode{}
	node := root
	for _, product := range products {
		node = root
		for _, c := range product {
			ind := c - 'a'
			if node.Children[ind] == nil {
				node.Children[ind] = &TrieNode{Val: string(c)}
				node.Children[ind].Children = [26]*TrieNode{}
			}
			node = node.Children[ind]
		}
		node.EndNode = true
		node.Val = product
	}
	res := [][]string{}
	for _, c := range searchWord {
		tmp := []string{}
		root = dfs1268(root, c, &tmp)
		res = append(res, tmp)
	}
	return res
}

func dfs1268(root *TrieNode, char rune, res *[]string) *TrieNode {
	if root == nil {
		return root
	}
	if len(*res) >= 3 {
		return root
	}
	if char != ' ' {
		if root.Children[char-'a'] == nil {
			return nil
		}
		root = root.Children[char-'a']
	}
	if root.EndNode {
		*res = append(*res, root.Val)
	}
	for _, child := range root.Children {
		if child == nil {
			continue
		}
		dfs1268(child, ' ', res)
	}
	return root
}
