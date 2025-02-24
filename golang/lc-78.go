package main

func subsets(nums []int) [][]int {
	ans := [][]int{}
	dfs78(nums, 0, []int{}, &ans)
	return ans
}

func dfs78(nums []int, curPos int, curVal []int, ans *[][]int) {
	if curPos == len(nums) {
		*ans = append(*ans, curVal)
		return
	}
	dfs78(nums, curPos+1, curVal, ans)
	dfs78(nums, curPos+1, append(curVal, nums[curPos]), ans)
}
