package main

func bsearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		ind := bsearch(numbers[i+1:], target-numbers[i])
		if ind != -1 {
			return []int{i + 1, i + 1 + ind + 1}
		}
	}
	return []int{}
}
