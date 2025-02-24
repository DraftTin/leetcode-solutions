package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxVal := 0
	for l < r {
		if height[l] < height[r] {
			maxVal = max(maxVal, (r-l)*height[l])
			l++
		} else {
			maxVal = max(maxVal, (r-l)*height[r])
			r--
		}
	}
	return maxVal
}
