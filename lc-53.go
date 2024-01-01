func maxSubArray(nums []int) int {
	ans := math.MinInt32
	cur := 0
	for _, num := range nums {
		if cur <= 0 {
			cur = num
		} else {
			cur += num
		}
		ans = max(ans, cur)
	}
	return ans
}
