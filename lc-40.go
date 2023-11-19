func combinationSum2(candidates []int, target int) [][]int {
	curCombination := make([]int, len(candidates))
	visited := make([]int, len(candidates))
	ans := [][]int{}
	helper(candidates, 0, &curCombination, 0, &ans, target, visited)
	return ans
}

func helper(candidates []int, curSum int, curCombination *[]int, curPos int, ans *[][]int, target int, visited []bool) {
	if curSum == target {
		item := make([]int, curPos)
		copy(item, curCombination)
		*ans = append(*ans, item)
		return
	}
	for i, _ := range candidates {
		if visited[i] {
			continue
		}
		if candidates[i]+curSum > target {
			continue
		}
		visited[i] = true
		(*curCombination)[curPos] = candidates[i]
		helper(candidates, curSum+candidates[i], curCombination, curPos+1, ans, target, visited)
		visited[i] = false
	}
}
