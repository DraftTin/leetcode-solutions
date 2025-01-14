package main

// Recursive DFS: add flag -> recurse(dfs) -> restore the flag
// try from big value to small value can guarantee the maximum result
func constructDistancedSequence(n int) []int {
	res := make([]int, 2*n-1)
	flag := make([]bool, n+1)
	dfs(res, flag, 0)
	return res
}

func dfs(arr []int, flag []bool, curPos int) bool {
	if curPos == len(arr) {
		return true
	}
	if arr[curPos] != 0 {
		return dfs(arr, flag, curPos+1)
	}
	for i := len(flag) - 1; i >= 1; i-- {
		if flag[i] {
			continue
		}
		if i == 1 {
			flag[i], arr[curPos] = true, 1
			if dfs(arr, flag, curPos+1) {
				return true
			}
			flag[i], arr[curPos] = false, 0
		} else if curPos+i < len(arr) && arr[curPos+i] == 0 {
			arr[curPos], arr[curPos+i], flag[i] = i, i, true
			if dfs(arr, flag, curPos+1) {
				return true
			}
			arr[curPos], arr[curPos+i], flag[i] = 0, 0, false
		}
	}
	return false
}
