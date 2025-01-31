package main

func minimumOperations(nums []int) int {
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		if nums[left] < nums[right] {
			nums[left+1] += nums[left]
			left++
			count++
		} else if nums[left] > nums[right] {
			nums[right-1] += nums[right]
			right--
			count++
		} else {
			left++
			right--
		}
	}
	return count
}
