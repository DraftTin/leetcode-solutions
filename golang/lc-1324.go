import (
    "fmt"
    "strings"
)

func printVertically(s string) []string {
  wordList := strings.Split(s, " ")
  maxLen := 0
  for _, word := range wordList {
    maxLen = max(maxLen, len(word))
  }
  fmt.Println(wordList)
  ans := make([]string, 0, maxLen)
  n := len(wordList)
  for i := 0; i < maxLen; i++ {
    tmp := make([]uint8, 0, n)
    for j := 0; j < n; j++ {
      if len(wordList[j]) > i {
        tmp = append(tmp, wordList[j][i])
      } else {
        tmp = append(tmp, ' ')
      }
    }
    for j := n - 1; j >= 0 && tmp[j] == ' '; j-- {
      tmp = tmp[:len(tmp) - 1]
    }
    ans = append(ans, string(tmp))
  }
  return ans
}
