func maximum69Number (num int) int {
  s := []uint8(strconv.Itoa(num))
  for i := 0; i < len(s); i++ {
    if s[i] == '6' {
      s[i] = '9'
      break
    }
  }
  ans, _ := strconv.Atoi(string(s))

  return ans
}
