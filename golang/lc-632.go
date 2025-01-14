package main

import (
	"math"
)

// Merge K Arrays + Sliding Window
func smallestRange(nums [][]int) []int {
	totalSize := 0
	for _, arr := range nums {
		totalSize += len(arr)
	}
	arr := make([][2]int, totalSize)
	indices := make([]int, len(nums))
	for i := 0; i < totalSize; i++ {
		minVal := math.MaxInt32
		minIndex := 0
		for j := 0; j < len(indices); j++ {
			if indices[j] < len(nums[j]) && nums[j][indices[j]] < minVal {
				minVal = nums[j][indices[j]]
				minIndex = j
			}
		}
		arr[i] = [2]int{minVal, minIndex}
		indices[minIndex]++
	}
	windowMap := make(map[int]int)
	left := 0
	count := 0
	res := []int{-100000, 100001}
	for i := 0; i < len(arr); i++ {
		windowMap[arr[i][1]]++
		if windowMap[arr[i][1]] == 1 {
			count++
		} else {
			for left < i && windowMap[arr[left][1]] != 1 {
				windowMap[arr[left][1]]--
				left++
			}
		}
		if count == len(nums) && arr[i][0]-arr[left][0] < res[1]-res[0] {
			res[0] = arr[left][0]
			res[1] = arr[i][0]
		}
	}
	return res
}
