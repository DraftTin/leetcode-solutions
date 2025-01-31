package main

// Sliding Window
func minSwaps(data []int) int {
	totalCount := 0
	for _, num := range data {
		totalCount += num
	}
	left, right := 0, 0
	curCount := 0
	maxCount := 0
	for right < len(data) {
		curCount += data[right]
		maxCount = max(maxCount, curCount)
		if right-left+1 == totalCount {
			curCount -= data[left]
			left++
		}
		right++
	}
	return totalCount - maxCount
}
