func minimizeSum(nums []int) int {
  max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
  min1, min2, min3 := math.MaxInt32, math.MaxInt32, math.MaxInt32
  for _, num := range nums {
    if num > max1 {
      max3 = max2
      max2 = max1
      max1 = num
    } else if num > max2 {
      max3 = max2
      max2 = num
    } else if num > max3 {
      max3 = num
    }

    if num < min1 {
      min3 = min2
      min2 = min1
      min1 = num
    } else if num < min2 {
      min3 = min2
      min2 = num
    } else if num < min3 {
      min3 = num
    }
  }
  return min(min(max1 - min3, max3 - min1), max2 - min2)
}

