func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	if i == len(intervals) {
		intervals = append(intervals, newInterval)
		return intervals
	}
	if newInterval[1] < intervals[i][0] {
		return append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
	}
	rbound := max(intervals[i][1], newInterval[1])
	intervals[i][1] = rbound
	intervals[i][0] = min(intervals[i][0], newInterval[0])
	j := i + 1
	for j < len(intervals) {
		if intervals[j][1] > rbound {
			if intervals[j][0] > rbound {
				intervals = append(intervals[:i+1], intervals[j:]...)
				return intervals
			} else {
				rbound = intervals[j][1]
				intervals[i][1] = rbound
			}
		}
		j++
	}
	return intervals[:i+1]
}
