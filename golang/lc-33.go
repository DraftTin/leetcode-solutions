func search(nums []int, target int) int {
  n := len(nums)
  left, right := 0, n - 1
  var pos int
  if nums[0] >= nums[n - 1] {
    l, r := 0, n - 1
    for l <= r {
      mid := (l + r) / 2
      if nums[mid] >= nums[0] {
        if (mid + 1 == n || nums[mid + 1] < nums[0]) {
            pos = mid
            break
        } else {
          l = mid + 1
        }
      } else {
        r = mid - 1
      }
    }
    if target >= nums[0] && target <= nums[pos] {
      right = pos
    } else {
      left = pos + 1
    }
  }
  newNums := nums[left:right + 1]
  index := sort.Search(len(newNums), func(i int)bool {
    return newNums[i] >= target
  })
  if index != len(newNums) && newNums[index] == target {
    return left + index
  }
  return -1
}
