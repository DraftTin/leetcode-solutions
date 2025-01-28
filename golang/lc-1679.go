package main

import "fmt"

func maxOperations(nums []int, k int) int {
	count := map[int]int{}
	res := 0
	for _, num := range nums {
		count[num]++
		fmt.Println(num, count[num])
		if count[k-num] > 0 {
			count[k-num]--
			count[num]--
			res++
		}
	}
	return res
}
