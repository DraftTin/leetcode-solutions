func search(nums []int, target int) int {
  index := sort.Search(len(nums), func(i int) bool {
    return nums[i] >= target
  })
  if index != len(nums) && nums[index] == target {
    return index
  }
  return -1
}
