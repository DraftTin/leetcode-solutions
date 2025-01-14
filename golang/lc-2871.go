func maxSubarrays(nums []int) int {
	v := -1
	ans := 0
	for _, num := range nums {
		v &= num
		if v == 0 {
			v = -1
			ans++
		}
	}
	return max(ans, 1)
}
