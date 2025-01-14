package main

func maxSum(nums []int, m int, k int) int64 {
	sumVal := 0
	maxSumVal := int64(0)
	cur := 0
	start := 0
	table := map[int]int{}
	count := 0
	for cur < len(nums) {
		sumVal += nums[cur]
		if table[nums[cur]] == 0 {
			count++
		}
		table[nums[cur]]++
		if cur-start == k {
			sumVal -= nums[start]
			if table[nums[start]] == 1 {
				count--
			}
			table[nums[start]]--
			start++
		}
		if count >= m {
			maxSumVal = max(maxSumVal, int64(sumVal))
		}
		cur++
	}
	return maxSumVal
}
