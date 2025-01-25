package main

func eventualSafeNodes(graph [][]int) []int {
	res := make(map[int]bool)
	visited := make([]bool, len(graph))
	resVal := []int{}
	for i := 0; i < len(graph); i++ {
		isSafe := dfs802(graph, i, &visited, res)
		if isSafe {
			resVal = append(resVal, i)
		}
	}
	return resVal
}

func dfs802(graph [][]int, index int, visited *[]bool, res map[int]bool) bool {
	if val, ok := res[index]; ok == true {
		return val
	}
	if (*visited)[index] == true {
		res[index] = false
		return false
	}
	(*visited)[index] = true
	for _, node := range graph[index] {
		isSafe := dfs802(graph, node, visited, res)
		if isSafe == false {
			res[index] = false
			return false
		}
	}
	res[index] = true
	return true
}
