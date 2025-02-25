package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 0 unvisited 1 ok - no cycle, 2 - visiting
	visited := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if dfs207(graph, visited, i) == false {
				return false
			}
		}
	}
	return true
}

func dfs207(graph [][]int, visited []int, curPos int) bool {
	visited[curPos] = 2
	for _, node := range graph[curPos] {
		if visited[node] == 2 {
			return false
		}
		if visited[node] == 0 {
			res := dfs207(graph, visited, node)
			if res == false {
				return false
			}
		}
	}
	visited[curPos] = 1
	return true
}
