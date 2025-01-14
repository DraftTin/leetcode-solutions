func getLastMoment(n int, left []int, right []int) int {
  ans := 0
  for _, pos := range right {
    if n - pos > ans {
      ans = n - pos
    }
  }
  for _, pos := range left {
    if pos > ans {
      ans = pos
    }
  }
  return ans
}
