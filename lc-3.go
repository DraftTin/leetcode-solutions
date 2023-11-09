func lengthOfLongestSubstring(s string) int {
  mp := map[rune]int{}
  ans := 0
  left := 0
  for i, c := range s {
    if pos, ok := mp[c]; ok == true && pos >= left {
      left = pos + 1
    }  
    mp[c] = i
    ans = max(ans, i - left + 1)
  }
  return ans
}
