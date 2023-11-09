func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
  if fx == sx && fy == sy {
    return t != 1
  }
  rows := int(math.Abs(float64(fy - sy)))
  cols := int(math.Abs(float64(fx - sx))) 
  distance := max(rows, cols)
  return distance <= t
}

