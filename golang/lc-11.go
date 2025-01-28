package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := (r - l) * min(height[l], height[r])
	for l < r {
		for l < r && height[l] <= height[r] {
			l++
			res = max(res, (r-l)*min(height[l], height[r]))
		}
		for l < r && height[r] <= height[l] {
			r--
			res = max(res, (r-l)*min(height[l], height[r]))
		}
	}
	return res
}
