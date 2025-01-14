package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []int64

func (a PriorityQueue) Len() int           { return len(a) }
func (a PriorityQueue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriorityQueue) Less(i, j int) bool { return a[i] >= a[j] }

func (a *PriorityQueue) Push(val any) {
	*a = append(*a, val.(int64))
}

func (a *PriorityQueue) Pop() any {
	old := *a
	n := len(old)
	res := old[n-1]
	*a = old[:n-1]
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// PriorityQueue + BFS
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	pq := PriorityQueue{}
	heap.Init(&pq)
	que := []*TreeNode{root}
	size := 1
	levelSum := int64(0)
	level := 0
	for len(que) != 0 {
		tmp := que[0]
		levelSum += int64(tmp.Val)
		que = que[1:]
		if tmp.Left != nil {
			que = append(que, tmp.Left)
		}
		if tmp.Right != nil {
			que = append(que, tmp.Right)
		}
		size--
		if size == 0 {
			size = len(que)
			fmt.Println(levelSum)
			heap.Push(&pq, levelSum)
			levelSum = 0
			level++
		}
	}
	if level < k {
		return -1
	}
	res := int64(0)
	for i := 0; i < k; i++ {
		res = heap.Pop(&pq).(int64)
	}
	return res
}
