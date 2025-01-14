package main

import "sort"

// Dynamic Programming
// p(i): largest index j, i < j such that job i and job j is compatible
// dp[i]: the maximum profit found until index i
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	intervals := make([][3]int, len(startTime))
	for i := 0; i < len(intervals); i++ {
		intervals[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	dpTable := make([]int, len(intervals))
	dpTable[0] = intervals[0][2]
	for i := 1; i < len(intervals); i++ {
		l, r := 0, i-1
		for l <= r {
			mid := (l + r) / 2
			if intervals[mid][1] <= intervals[i][0] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if r == -1 {
			dpTable[i] = max(dpTable[i-1], intervals[i][2])
		} else {
			dpTable[i] = max(dpTable[i-1], intervals[i][2]+dpTable[r])
		}
	}
	return dpTable[len(intervals)-1]
}
