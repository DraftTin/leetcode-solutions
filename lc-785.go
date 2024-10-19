package main

func isBipartite(graph [][]int) bool {
	que := make([]int, 0)
	visited := make([]bool, len(graph))
	for i := 0; i < len(graph); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		redSet := map[int]bool{}
		blueSet := map[int]bool{}
		que = append(que, i)
		size := 1
		curSet := redSet
		nextSet := blueSet
		for len(que) != 0 {
			tmp := que[0]
			if nextSet[tmp] {
				return false
			}
			curSet[tmp] = true
			for j := 0; j < len(graph[tmp]); j++ {
				if curSet[graph[tmp][j]] {
					return false
				}
				if visited[graph[tmp][j]] == false {
					visited[graph[tmp][j]] = true
					que = append(que, graph[tmp][j])
				}
			}
			que = que[1:]
			size--
			if size == 0 {
				size = len(que)
				curSet, nextSet = nextSet, curSet
			}
		}
	}
	return true
}
