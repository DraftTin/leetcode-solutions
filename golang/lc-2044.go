package main

// Backtracking: there are totally 2^n subsets. Use DFS to visit every possibility
func countMaxOrSubsets(nums []int) int {
	maxOR := 0
	for _, num := range nums {
		maxOR |= num
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		count += dfs2044(nums, i, nums[i], maxOR)
	}
	return count
}

func dfs2044(nums []int, curPos, curOR, maxOR int) int {
	count := 0
	if curOR == maxOR {
		count++
	}
	for i := curPos + 1; i < len(nums); i++ {
		count += dfs2044(nums, i, curOR|nums[i], maxOR)
	}
	return count
}
