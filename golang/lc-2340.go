package main

import "math"

func minimumSwaps(nums []int) int {
	lr, rl := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	maxVal, minVal := math.MinInt32, math.MaxInt32
	for i, num := range nums {
		if num > maxVal {
			maxVal = num
			lr = i
		} else if num == maxVal {
			lr = max(lr, i)
		}
		if num < minVal {
			minVal = num
			rl = i
		} else if num == minVal {
			rl = min(rl, i)
		}
	}
	res := len(nums) - 1 - lr + rl
	if lr < rl {
		res--
	}
	return res
}
