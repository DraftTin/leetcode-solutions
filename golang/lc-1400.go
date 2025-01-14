/*
Time Complexity: O(n)
Space Complexity O(n)

Solution: For each palindrome, we can assign up to one asymmeric character. Therefore,
If the number of single characters are greater than k, then return false.
*/

func canConstruct(s string, k int) bool {
  if len(s) < k {
    return false
  }
  mp := map[rune]int{}
  for _, c := range s {
    mp[c]++
  }
  for _, val := range mp {
    if val % 2 == 1 {
      k--
    }
  }
  if k < 0 {
    return false
  }
  return true 
}

