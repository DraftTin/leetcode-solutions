package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		indegrees[edge[1]]++
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	que := []int{}
	for index, val := range indegrees {
		if val == 0 {
			que = append(que, index)
		}
	}
	n := 0
	for len(que) != 0 {
		course := que[0]
		que = que[1:]
		n++
		for _, c := range graph[course] {
			indegrees[c]--
			if indegrees[c] == 0 {
				que = append(que, c)
			}
		}
	}
	if n == numCourses {
		return true
	}
	return false
}
