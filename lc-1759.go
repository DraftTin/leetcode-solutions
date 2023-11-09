func countHomogenous(s string) int {
  mp := map[int]int{}
  i := 0
  n := len(s)
  for i < n {
    c := s[i]
    j := i
    for j < n && s[j] == c {
      j++
    }
    mp[j - i]++
    i = j
  }
  fmt.Println(mp)
  ans := 0
  for key, val := range mp {
    ans += ((key + 1) * key) / 2 * val
  }
  return ans
}
