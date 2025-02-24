package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}
	dfs39(candidates, 0, target, &ans, &[]int{})
	return ans
}

func dfs39(candidates []int, curPos, target int, ans *[][]int, curAns *[]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, (*curAns)...))
		return
	}
	if curPos == len(candidates) {
		return
	}
	if candidates[curPos] > target {
		return
	}
	dfs39(candidates, curPos+1, target, ans, curAns)
	*curAns = append(*curAns, candidates[curPos])
	dfs39(candidates, curPos, target-candidates[curPos], ans, curAns)
	*curAns = (*curAns)[:len(*curAns)-1]
}
