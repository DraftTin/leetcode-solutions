func combinationSum2(candidates []int, target int) [][]int {
	curCombination := make([]string, len(candidates))
	visited := make([]bool, len(candidates))
	sort.Ints(candidates)
	ans := map[string]bool{}
	helper(candidates, 0, 0, &curCombination, 0, ans, target, visited)
	result := make([][]int, 0, len(ans))
	for key, _ := range ans {
		item := strings.Split(key, ",")
		numArr := make([]int, len(item))
		for i := 0; i < len(item); i++ {
			numArr[i], _ = strconv.Atoi(item[i])
		}
		result = append(result, numArr)
	}
	return result
}

func helper(candidates []int, pos int, curSum int, curCombination *[]string, curPos int, ans map[string]bool, target int, visited []bool) {
	if curSum == target {
		item := make([]string, curPos)
		copy(item, *curCombination)
		key := strings.Join(item, ",")
		ans[key] = true
		return
	}
	for i := pos; i < len(candidates); i++ {
		if visited[i] {
			continue
		}
		if candidates[i]+curSum > target {
			continue
		}
		visited[i] = true
		(*curCombination)[curPos] = strconv.Itoa(candidates[i])
		helper(candidates, i+1, curSum+candidates[i], curCombination, curPos+1, ans, target, visited)
		visited[i] = false
	}
}
