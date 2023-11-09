func minWindow(s string, t string) string {
  counts := map[uint8]int{}
  for i, _ := range t {
    counts[t[i]]++
  }
  start, end := 0, 0
  window := map[uint8]int{}
  ans := [2]int{0, math.MaxInt32}
  cur := 0
  for end < len(s) {
    if counts[s[end]] > window[s[end]] {
      cur++
    }
    window[s[end]]++
    for start <= end && window[s[start]] > counts[s[start]] {
      window[s[start]]--
      start++
    }
    if cur == len(t) && end - start + 1 < ans[1] - ans[0] + 1{
      ans[0], ans[1] = start, end
    }
    end++
  }
  if ans[1] == math.MaxInt32 {
    return ""
  }
  return s[ans[0] : ans[1] + 1]
}
