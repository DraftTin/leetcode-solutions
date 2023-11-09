func maxProfit(prices []int) int {
  minVal := math.MaxInt32
  ans := 0
  for _, val := range prices {
    minVal = min(minVal, val)
    ans = max(val - minVal, ans)
  }
  return ans
}
