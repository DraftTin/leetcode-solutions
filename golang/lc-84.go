package main

func largestRectangleArea(heights []int) int {
	left := make([]int, len(heights))
	right := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		j := i - 1
		for j >= 0 && heights[j] >= heights[i] {
			j = left[j] - 1
		}
		left[i] = j + 1
	}
	for i := len(heights) - 1; i >= 0; i-- {
		j := i + 1
		for j < len(heights) && heights[j] >= heights[i] {
			j = right[j] + 1
		}
		right[i] = j - 1
	}
	res := 0
	for i := 0; i < len(heights); i++ {
		res = max(res, heights[i]*(right[i]-left[i]+1))
	}
	return res
}
