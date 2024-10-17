package main

func minLength(s string) int {
	visited := make([]bool, len(s))
	i := 0
	for i < len(s) {
		l, r := i, i+1
		for l >= 0 && r < len(s) && (s[l] == 'A' && s[r] == 'B' || s[l] == 'C' && s[r] == 'D') {
			visited[l] = true
			visited[r] = true
			for l >= 0 && visited[l] == true {
				l--
			}
			for r <= len(s)-1 && visited[r] == true {
				r++
			}
		}
		i = r
	}
	res := 0
	for _, val := range visited {
		if val == true {
			res++
		}
	}
	return res
}
