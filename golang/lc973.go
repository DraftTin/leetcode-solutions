package main

import "container/heap"

type Heap973 [][]int

func (this Heap973) Len() int {
	return len(this)
}

func (this Heap973) Less(i, j int) bool {
	return this[i][0]*this[i][0]+this[i][1]*this[i][1] < this[j][0]*this[j][0]+this[j][1]*this[j][1]
}

func (this Heap973) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Heap973) Pop() interface{} {
	old := *this
	n := len(*this)
	res := old[n-1]
	*this = old[:n-1]
	return res
}

func (this *Heap973) Push(item interface{}) {
	*this = append(*this, item.([]int))
}

func kClosest(points [][]int, k int) [][]int {
	pque := Heap973{}
	heap.Init(&pque)
	for _, point := range points {
		heap.Push(&pque, point)
	}
	res := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		point := heap.Pop(&pque)
		res = append(res, point.([]int))
	}
	return res
}
