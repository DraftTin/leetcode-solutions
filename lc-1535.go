func getWinner(arr []int, k int) int {
  n := len(arr)
  if k > n {
    maxVal := math.MinInt32
    for _, num := range arr {
      maxVal = max(maxVal, num)
    }
    return maxVal
  }
  q := make([]int, n)
  for i, num := range arr {
    q[i] = num
  }
  count := 0
  for {
    if q[0] > q[1] {
      q[0], q[1] = q[1], q[0]
      q = append(q[1:], q[0])
      count++
    } else {
      q = append(q[1:], q[0])
      count = 1
    }
    if count == k {
      return q[0]
    }
  }
  return 0
}
