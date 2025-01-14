package main

import (
	"container/heap"
	"sort"
)

type PriorQueue []int

func (a PriorQueue) Len() int           { return len(a) }
func (a PriorQueue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriorQueue) Less(i, j int) bool { return a[i] < a[j] }

func (a *PriorQueue) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *PriorQueue) Pop() any {
	old := *a
	n := len(old)
	res := old[n-1]
	*a = old[:n-1]
	return res
}

// Using Prior Queue to output the smalest number
// note that when sort those operations, we need to put 'leaving' friends before the 'coming' friends so that seats can be released correctly.
func smallestChair(times [][]int, targetFriend int) int {
	list := make([][3]int, 0, 2*len(times))
	for i := 0; i < len(times); i++ {
		list = append(list, [3]int{times[i][0], 0, i})
		list = append(list, [3]int{times[i][1], 1, i})
	}
	arriveTime := times[targetFriend][0]
	sort.Slice(list, func(i, j int) bool {
		if list[i][0] < list[j][0] {
			return true
		} else if list[i][0] == list[j][0] {
			return list[i][1] == 1
		}
		return false
	})

	maxNumber := 0
	friendMap := make(map[int]int)
	que := &PriorQueue{}
	heap.Init(que)
	for i := 0; i < len(list); i++ {
		if list[i][1] == 0 {
			if que.Len() == 0 {
				friendMap[list[i][2]] = maxNumber
				maxNumber++
			} else {
				num := heap.Pop(que)
				friendMap[list[i][2]] = num.(int)
			}
			if arriveTime == list[i][0] {
				return friendMap[list[i][2]]
			}
		} else {
			heap.Push(que, friendMap[list[i][2]])
		}
	}
	return 0
}
