func numTeams(rating []int) int {
  ans := 0
  n := len(rating)
  for i := 1; i < n; i++ {
    less, greater := [2]int{}, [2]int{}
    for j := 0; j < n; j++ {
      index := 0
      if j > i {
        index = 1
      }
      if rating[j] < rating[i] {
        less[index]++
      }
      if rating[j] > rating[i] {
        greater[index]++
      }
    }
    ans += less[0] * greater[1] + less[1] * greater[0]
  }
  return ans
}
