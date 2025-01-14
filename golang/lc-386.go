package main

// DFS
func lexicalOrder(n int) []int {
	res := make([]int, n)
	dfsForlexicalOrder(res, 0, 1)
	return res
}

func dfsForlexicalOrder(arr []int, curPos, curNum int) int {
	n := len(arr)
	if curNum > n {
		return curPos
	}
	arr[curPos] = curNum
	curPos = dfsForlexicalOrder(arr, curPos+1, curNum*10)
	if curNum%10 == 9 {
		return curPos
	}
	return dfsForlexicalOrder(arr, curPos, curNum+1)
}
