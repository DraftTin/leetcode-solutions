func findDuplicate(nums []int) int {
  fast, slow := 0, 0
  for {
    fast = nums[nums[fast]]
    slow = nums[slow]
    if fast == slow {
      break
    }
  }
  fast = 0
  for fast ! slow {
    fast, slow = nums[fast], nums[slow]
  }
  return fast
}
