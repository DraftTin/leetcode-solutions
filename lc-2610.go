package main

func findMatrix(nums []int) [][]int {
	visited := make([]bool, len(nums))
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			numMap := make(map[int]bool)
			ans = append(ans, helper(nums, visited, i, numMap))
		}
	}
	return ans
}

func helper(nums []int, visited []bool, pos int, numMap map[int]bool) []int {
	if pos == len(nums) {
		return []int{}
	}
	visited[pos] = true
	numMap[nums[pos]] = true
	for i := pos + 1; i < len(nums); i++ {
		if !visited[i] && numMap[nums[i]] == false {
			return append(helper(nums, visited, i, numMap), nums[pos])
		}
	}
	return []int{nums[pos]}
}
