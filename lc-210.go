package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegrees := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		indegrees[edge[0]]++
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	que := []int{}
	for index, val := range indegrees {
		if val == 0 {
			que = append(que, index)
		}
	}
	outputOrder := make([]int, 0, numCourses)
	for len(que) != 0 {
		course := que[0]
		que = que[1:]
		outputOrder = append(outputOrder, course)
		for _, c := range graph[course] {
			indegrees[c]--
			if indegrees[c] == 0 {
				que = append(que, c)
			}
		}
	}
	if len(outputOrder) != numCourses {
		return []int{}
	}
	return outputOrder
}
