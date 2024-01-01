func findTargetSumWays(nums []int, target int) int {
	dp := make(map[[2]int]int)
	return dfs(nums, 0, target, dp)
}

func dfs(nums []int, pos int, target int, dp map[[2]int]int) int {
	if val, ok := dp[[2]int{target, pos}]; ok == true {
		return val
	}
	if pos == len(nums) {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	count := 0
	count += dfs(nums, pos+1, target-nums[pos], dp)
	count += dfs(nums, pos+1, target+nums[pos], dp)
	dp[[2]int{target, pos}] = count
	return count
}
