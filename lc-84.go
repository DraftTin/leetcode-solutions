func largestRectangleArea(heights []int) int {
  n := len(heights)
  left := make([]int, n)
  right := make([]int, n)
  for i := 0; i < n; i++ {
    j := i - 1
    for j >= 0 && heights[j] >= heights[i] {
      j = left[j] - 1
    }
    left[i] = j + 1
  }
  for i := n - 1; i >= 0; i-- {
    j := i + 1
    for j < n && heights[j] >= heights[i] {
      j = right[j] + 1
    }
    right[i] = j - 1
  }
  ans := 0
  for i := 0; i < n; i++ {
    ans = max(ans, (right[i] - left[i] + 1) * heights[i])
  }
  return ans
}
