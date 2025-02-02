package main

func maximumSubarraySum(nums []int, k int) int64 {
	l, r := 0, 0
	count := map[int]int{}
	res := int64(0)
	curSum := int64(0)
	for r < len(nums) {
		count[nums[r]]++
		curSum += int64(nums[r])
		if r-l+1 > k {
			count[nums[l]]--
			curSum -= int64(nums[l])
			l++
		}
		for count[nums[r]] > 1 {
			count[nums[l]]--
			curSum -= int64(nums[l])
			l++
		}
		if r-l+1 == k && curSum > res {
			res = curSum
		}
		r++
	}
	return res
}
