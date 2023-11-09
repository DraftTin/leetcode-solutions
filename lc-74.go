func searchMatrix(matrix [][]int, target int) bool {
  n, m := len(matrix), len(matrix[0])
  i, j := 0, n * m - 1
  for i <= j {
    mid := (i + j) / 2
    row, col := mid / m, mid % m
    if matrix[row][col] == target {
      return true
    } else if matrix[row][col] < target {
      i = mid + 1
    } else {
      j = mid - 1
    }
  }
  return false
}


