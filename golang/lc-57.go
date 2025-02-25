package main

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	for i < len(intervals) && intervals[i][0] <= newInterval[0] {
		i++
	}
	if i != 0 {
		linterval := intervals[i-1]
		if linterval[1] >= newInterval[0] {
			linterval[1] = max(linterval[1], newInterval[1])
			i--
		} else {
			tmp := append(append([][]int{}, intervals[:i]...), newInterval)
			intervals = append(tmp, intervals[i:]...)
		}
	} else {
		intervals = append(append([][]int{}, newInterval), intervals...)
	}
	j := i + 1
	for j < len(intervals) && intervals[j][0] <= intervals[i][1] {
		intervals[i][1] = max(intervals[i][1], intervals[j][1])
		j++
	}
	return append(intervals[:i+1], intervals[j:]...)
}
