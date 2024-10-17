package main

import "math"

// Prefix Sum and Hash Map Optimization
func minSubarray(nums []int, p int) int {
	remainderMap := make(map[int]int)
	sumVal := 0
	for _, num := range nums {
		sumVal += num
	}
	if sumVal%p == 0 {
		return 0
	}
	cur := 0
	res := math.MaxInt32
	for i, num := range nums {
		cur += num
		remainderMap[cur%p] = i
		rightRemainder := (sumVal - cur) % p
		targetRemainder := (p - rightRemainder) % p
		if index, ok := remainderMap[targetRemainder]; ok {
			res = min(res, i-index)
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
