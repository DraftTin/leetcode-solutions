package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	res := 1
	top := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[top][1] {
			top = i
			res++
		} else if intervals[i][1] < intervals[top][1] {
			top = i
		}
	}
	return res
}
