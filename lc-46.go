func permute(nums []int) [][]int {
	ans := [][]int{}
	permutation := make([]int, len(nums))
	visited := make([]bool, len(nums))
	dfs(nums, &ans, permutation, 0, visited)
	return ans

}

func dfs(nums []int, ans *[][]int, combination []int, curPos int, visited []bool) {
	if curPos == len(nums) {
		item := make([]int, len(nums))
		copy(item, combination)
		*ans = append(*ans, item)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		combination[curPos] = nums[i]
		dfs(nums, ans, combination, curPos+1, visited)
		visited[i] = false
	}
}
