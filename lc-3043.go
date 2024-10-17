package main

import "strconv"

type TrieNode struct {
	nodes []*TrieNode
	flag  int
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	root := &TrieNode{nodes: make([]*TrieNode, 10)}
	str1, str2 := make([]string, 0, len(arr1)), make([]string, 0, len(arr2))
	for i := 0; i < len(arr1); i++ {
		str1 = append(str1, strconv.Itoa(arr1[i]))
	}
	for i := 0; i < len(arr2); i++ {
		str2 = append(str2, strconv.Itoa(arr2[i]))
	}
	for _, val := range str1 {
		cur := root
		for i := 0; i < len(val); i++ {
			if cur.nodes[val[i]-'0'] == nil {
				cur.nodes[val[i]-'0'] = &TrieNode{nodes: make([]*TrieNode, 10), flag: 1}
			}
			cur = cur.nodes[val[i]-'0']
		}
	}
	res := 0
	for _, val := range str2 {
		cur := root
		depth := 0
		for i := 0; i < len(val); i++ {
			depth++
			if cur.nodes[val[i]-'0'] == nil {
				cur.nodes[val[i]-'0'] = &TrieNode{nodes: make([]*TrieNode, 10), flag: 2}
			}
			cur = cur.nodes[val[i]-'0']
			if cur.flag == 1 {
				res = max(res, depth)
			}
		}
	}
	return res
}
