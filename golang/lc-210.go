package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	incount := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		incount[prerequisite[0]]++
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	que := []int{}
	n := 0
	for i, count := range incount {
		if count == 0 {
			n++
			que = append(que, i)
		}
	}
	res := []int{}
	for len(que) != 0 {
		tmp := que[0]
		que = que[1:]
		res = append(res, tmp)
		for _, course := range graph[tmp] {
			incount[course]--
			if incount[course] == 0 {
				n++
				que = append(que, course)
			}
		}
	}
	if n != numCourses {
		return []int{}
	}
	return res
}
