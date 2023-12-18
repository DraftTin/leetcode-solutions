func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	left, right := -1, -1
	ans := 0
	for _, interval := range intervals {
		if interval[0] > left && interval[1] > right {
			left = interval[0]
			ans++
		}
		right = max(right, interval[1])
	}
	return ans
}
