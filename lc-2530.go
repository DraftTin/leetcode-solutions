package main

import "container/heap"

type PriorityQue []int

func (a PriorityQue) Len() int           { return len(a) }
func (a PriorityQue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriorityQue) Less(i, j int) bool { return a[i] > a[j] }

func (a *PriorityQue) Pop() any {
	old := *a
	n := len(*a)
	res := old[n-1]
	*a = old[:n-1]
	return res
}

func (a *PriorityQue) Push(val any) {
	*a = append(*a, val.(int))
}

// PriorityQue
func maxKelements(nums []int, k int) int64 {
	que := make(PriorityQue, 0)
	heap.Init(&que)
	for _, num := range nums {
		heap.Push(&que, num)
	}
	score := int64(0)
	for i := 0; i < k; i++ {
		score += int64(que[0])
		tmp := que[0] % 3
		que[0] /= 3
		if tmp != 0 {
			que[0]++
		}
		heap.Fix(&que, 0)
	}
	return score
}
