package main

import (
	"math"
)

func shortestSubarray(nums []int, k int) int {
	left, right := 0, 0
	res := math.MaxInt32
	curSum := 0
	leftSum := make([]int, len(nums))
	totalSum := 0
	negArr := []int{}
	for right < len(nums) {
		totalSum += nums[right]
		leftSum[right] = totalSum
		curSum += nums[right]
		if nums[right] < 0 {
			negArr = append(negArr, right)
		}
		if curSum <= 0 {
			right++
			left = right
			curSum = 0
			negArr = []int{}
			continue
		}
		if curSum >= k {
			tmpInd := -1
			cutPos := -1
			for i, ind := range negArr {
				tmpSum := leftSum[ind]
				if left != 0 {
					tmpSum -= leftSum[left-1]
				}
				if curSum-tmpSum >= k {
					tmpInd = ind
					cutPos = i
				}
			}
			if tmpInd != -1 {
				left = tmpInd + 1
				curSum = leftSum[right] - leftSum[left-1]
				negArr = negArr[cutPos+1:]
			}
			for left < right && curSum-nums[left] >= k {
				curSum -= nums[left]
				left++
			}
			res = min(res, right-left+1)
			curSum = leftSum[right]
			if left != 0 {
				curSum -= leftSum[left-1]
			}
		}
		right++
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
