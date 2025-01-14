func characterReplacement(s string, k int) int {
  counts := [26]int{}
  start := 0
  maxCount := 0
  ans := 0
  for i := 0; i < len(s); i++ {
    counts[s[i] - 'A']++
    maxCount = max(maxCount, counts[s[i] - 'A']) 
    for i - start + 1 - maxCount > k {
      counts[s[start] - 'A']--
      start++
    }
    ans = max(i - start + 1, ans) 
  }
  return ans
}
