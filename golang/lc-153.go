func findMin(nums []int) int {
  if nums[0] < nums[len(nums) - 1] {
    return nums[0]
  }
  i, j := 0, len(nums) - 1
  for i < j {
    mid := (i + j) / 2
    if nums[mid] >= nums[0] {
      i = mid + 1
    } else if nums[mid] <= nums[len(nums) - 1] {
      j = mid
    }
  }
  return nums[j]
}
