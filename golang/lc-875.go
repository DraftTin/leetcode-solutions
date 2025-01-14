func minEatingSpeed(piles []int, h int) int {
  maxVal := math.MinInt32
  for _, val := range piles {
    maxVal = max(maxVal, val)
  }
  i, j := 1, maxVal
  for i < j {
    k := (i + j) / 2
    cost := 0
    for _, val := range piles {
      cost += val / k
      if val % k != 0 {
        cost += 1
      }
      if cost > h {
        break
      }
    }
    if cost > h {
      i = k + 1
    } else if cost <= h {
      j = k
    }
  }
  return j
}
