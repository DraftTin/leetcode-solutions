func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1][0], points[i-1][1]
		x2, y2 := points[i][0], points[i][1]
		ans += max(int(math.Abs(float64(x1-x2))), int(math.Abs(float64(y1-y2))))
	}
	return ans
}
