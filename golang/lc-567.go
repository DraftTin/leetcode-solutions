func checkInclusion(s1 string, s2 string) bool {
  counts := [26]int{}
  window := [26]int{}
  total := 0
  for _, c := range s1 {
    total++
    counts[c - 'a']++
  }
  start := 0
  end := 0
  for end < len(s2) {
    window[s2[end] - 'a']++
    for start <= end && window[s2[end] - 'a'] > counts[s2[end] - 'a'] {
        window[s2[start] - 'a']--
        start++
    }
    if end - start + 1 == total {
        return true
    }
    end++ 
  }
  return false
}
