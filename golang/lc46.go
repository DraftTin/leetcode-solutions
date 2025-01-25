package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	combination := make([]int, len(nums))
	visited := make([]bool, len(nums))
	dfs46(nums, 0, combination, &ans, visited)
	return ans
}

func dfs46(nums []int, curPos int, combination []int, ans *[][]int, visited []bool) {
	if curPos == len(nums) {
		newCombination := make([]int, len(nums))
		copy(newCombination, combination)
		*ans = append(*ans, newCombination)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		combination[curPos] = nums[i]
		dfs46(nums, curPos+1, combination, ans, visited)
		visited[i] = false
	}
}
