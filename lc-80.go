package main

func removeDuplicates(nums []int) int {
	ans := 0
	i := 0
	curPos := 0
	for i < len(nums) {
		count := 1
		j := i + 1
		for j < len(nums) {
			if nums[j] != nums[i] {
				break
			}
			j++
			count = 2
		}
		for k := 0; k < count; k++ {
			nums[curPos] = nums[i]
			curPos++
		}
		i = j
		ans += count
	}
	return ans
}
