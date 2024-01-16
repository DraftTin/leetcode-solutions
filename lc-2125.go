package main

func numberOfBeams(bank []string) int {
	n, m := len(bank), len(bank[0])
	pre := 0
	ans := 0
	for i := 0; i < n; i++ {
		cur := 0
		for j := 0; j < m; j++ {
			if bank[i][j] == '1' {
				cur++
			}
		}
		if cur == 0 {
			continue
		}
		ans += pre * cur
		pre = cur
	}
	return ans
}
