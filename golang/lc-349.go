package main

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}
	intersection := make(map[int]bool)
	for _, num := range nums2 {
		if set[num] == true {
			intersection[num] = true
		}
	}
	ans := make([]int, 0, len(intersection))
	for key := range intersection {
		ans = append(ans, key)
	}
	return ans
}
