func countNegatives(grid [][]int) int {
  count := 0
  n := len(grid)
  m := len(grid[0])
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      if grid[i][j] < 0 {
        count += m - j
        break
      }
    }
  }
  return count
}
