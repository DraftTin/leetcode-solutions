package main

import "sort"

// hashtable + digit manipulation
func maximumSum(nums []int) int {
	sort.Ints(nums)
	buckets := map[int][]int{}
	for i, num := range nums {
		tmp := 0
		for num != 0 {
			tmp += num % 10
			num /= 10
		}
		buckets[tmp] = append(buckets[tmp], nums[i])
	}
	res := -1
	for _, bucket := range buckets {
		if len(bucket) < 2 {
			continue
		}
		res = max(res, bucket[len(bucket)-1]+bucket[len(bucket)-2])
	}
	return res
}
