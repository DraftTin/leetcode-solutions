package main

import (
	"sort"
	"strings"
)

type TrieNode struct {
	SubNodes map[string]*TrieNode
	IsEnd    bool
	Val      string
}

// TrieTree
func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return len(folder[i]) < len(folder[j])
	})
	folderList := make([][]string, len(folder))
	root := &TrieNode{SubNodes: map[string]*TrieNode{}, IsEnd: false}
	marks := make([]bool, len(folder))
	for i := 0; i < len(folder); i++ {
		folderList[i] = strings.Split(folder[i], "/")
		curNode := root
		end := false
		for j := 1; j < len(folderList[i]); j++ {
			if curNode.IsEnd == true {
				end = true
			}
			if curNode.SubNodes[folderList[i][j]] == nil {
				if end == true {
					marks[i] = true
				}
				curNode.SubNodes[folderList[i][j]] = &TrieNode{SubNodes: map[string]*TrieNode{}, IsEnd: false, Val: folderList[i][j]}
				curNode = curNode.SubNodes[folderList[i][j]]
			} else {
				curNode = curNode.SubNodes[folderList[i][j]]
			}
		}
		curNode.IsEnd = true
	}
	res := []string{}
	for i := 0; i < len(folder); i++ {
		if marks[i] == false {
			res = append(res, folder[i])
		}
	}
	return res
}
