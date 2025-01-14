package main

type TrieNode2416 struct {
	Val  int
	Next [26]*TrieNode2416
}

// Trie Tree
func sumPrefixScores(words []string) []int {
	root := &TrieNode2416{}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			if node.Next[index] == nil {
				node.Next[index] = &TrieNode2416{Val: 0}
			}
			node = node.Next[index]
			node.Val++
		}
	}
	res := make([]int, 0, len(words))
	for _, word := range words {
		tmp := 0
		node := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			node = node.Next[index]
			tmp += node.Val
		}
		res = append(res, tmp)
	}
	return res
}
