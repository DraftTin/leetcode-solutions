func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	curCombination := make([]int, 150)
	if candidates[0] > target {
		return result
	}
	dfs(candidates, 0, 0, target, &curCombination, 0, &result)
	return result
}

func dfs(candidates []int, pos, curSum, target int, curCombination *[]int, curPos int, result *[][]int) {
	if curSum == target {
		combination := make([]int, curPos)
		copy(combination, *curCombination)
		*result = append(*result, combination)
		return
	}
	n := len(candidates)
	for i := pos; i < n; i++ {
		if curSum+candidates[i] > target {
			return
		}
		(*curCombination)[curPos] = candidates[i]
		dfs(candidates, i, curSum+candidates[i], target, curCombination, curPos+1, result)
	}
}
