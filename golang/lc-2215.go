package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	dict1, dict2 := map[int]bool{}, map[int]bool{}
	for _, num := range nums1 {
		dict1[num] = true
	}
	for _, num := range nums2 {
		dict2[num] = true
	}
	res := make([][]int, 2)
	for key, _ := range dict1 {
		if dict2[key] == false {
			res[0] = append(res[0], key)
		}
	}
	for key, _ := range dict2 {
		if dict1[key] == false {
			res[1] = append(res[1], key)
		}
	}
	return res
}
