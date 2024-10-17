package main

// Bit Operation
func xorAllNums(nums1 []int, nums2 []int) int {
	len1, len2 := len(nums1), len(nums2)
	res := 0
	if len1%2 != 0 {
		for _, num := range nums2 {
			res ^= num
		}
	}
	if len2%2 != 0 {
		for _, num := range nums1 {
			res ^= num
		}
	}
	return res
}
