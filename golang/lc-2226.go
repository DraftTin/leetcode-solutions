package main

func maximumCandies(candies []int, k int64) int {
	sum := 0
	for _, size := range candies {
		sum += size
	}
	minVal, maxVal := 1, sum/int(k)
	var ans int
	for minVal <= maxVal {
		midVal := (minVal + maxVal) / 2
		var n int64
		for _, size := range candies {
			n += int64(size / midVal)
		}
		if n >= k {
			ans = midVal
			minVal = midVal + 1
		} else {
			maxVal = midVal - 1
		}
	}
	return ans
}
