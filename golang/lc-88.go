package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}
	nums4 := nums1[i:]
	if i == m {
		nums4 = nums2[j:]
	}
	nums3 = append(nums3, nums4...)
	copy(nums1, nums3)
}
