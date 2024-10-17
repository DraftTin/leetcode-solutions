package main

import (
	"container/heap"
	"math"
)

type PriorQue []int

func (a PriorQue) Len() int           { return len(a) }
func (a PriorQue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriorQue) Less(i, j int) bool { return a[i] < a[j] }

func (a *PriorQue) Push(item any) {
	*a = append(*a, item.(int))
}

func (a *PriorQue) Pop() any {
	old := *a
	n := len(old)
	res := old[n-1]
	*a = old[:n-1]
	return res
}

func totalCost(costs []int, k int, candidates int) int64 {
	ll, lr, rl, rr := 0, candidates-1, len(costs)-candidates, len(costs)-1
	if lr >= rl {
		lr, rl = (len(costs)-1)/2, (len(costs)-1)/2+1
	}
	lque, rque := make(PriorQue, lr-ll+1), make(PriorQue, rr-rl+1)
	for i := ll; i <= lr; i++ {
		lque[i] = costs[i]
	}
	for i := rl; i <= rr; i++ {
		rque[i-rl] = costs[i]
	}
	heap.Init(&lque)
	heap.Init(&rque)
	var res int64
	for i := 0; i < k; i++ {
		lval, rval := math.MaxInt32, math.MaxInt32
		if len(lque) > 0 {
			lval = lque[0]
		}
		if len(rque) > 0 {
			rval = rque[0]
		}
		if rval < lval {
			heap.Pop(&rque)
			res += int64(rval)
			if rl-1 > lr {
				rl--
				heap.Push(&rque, costs[rl])
			}
		} else {
			heap.Pop(&lque)
			res += int64(lval)
			if rl-1 > lr {
				lr++
				heap.Push(&lque, costs[lr])
			}
		}
	}
	return res
}
