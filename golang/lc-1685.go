package main

import "math"

func getSumAbsoluteDifferences(nums []int) []int {
	res := make([]int, len(nums))
	baseVal := 0
	for _, num := range nums {
		baseVal += int(math.Abs(float64(num - nums[0])))
	}
	res[0] = baseVal
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		baseVal = baseVal - tmp - tmp*(len(nums)-1-i) + tmp*i
		res[i] = baseVal
	}
	return res
}
