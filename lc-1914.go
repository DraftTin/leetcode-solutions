func rotateGrid(grid [][]int, k int) [][]int {
	baseX, baseY := 0, 0
	m, n := len(grid), len(grid[0])
	rotate := make([][]int, m)
	for i := 0; i < m; i++ {
		rotate[i] = make([]int, n)
	}
	for m > 0 && n > 0 {
		circle := (m + n - 2) * 2
		move := k % circle
		for i := 0; i < circle; i++ {
			y1, x1 := getDest(m, n, i)
			y2, x2 := getDest(m, n, (i+move)%circle)
			rotate[y1+baseY][x1+baseX] = grid[baseY+y2][baseX+x2]
		}
		m -= 2
		n -= 2
		baseX += 1
		baseY += 1
	}
	return rotate
}

func getDest(m, n, step int) (int, int) {
	if step <= n-1 {
		return 0, step
	}
	if step <= m-1+n-1 {
		return step - (n - 1), n - 1
	}
	if step <= 2*(n-1)+(m-1) {
		return m - 1, (n - 1) - (step - (m - 1 + n - 1))
	} else {
		return (m+n-2)*2 - step, 0
	}
}

