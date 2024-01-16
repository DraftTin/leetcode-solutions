package main

func maxSlidingWindow(nums []int, k int) []int {
	que := make([]int, 0, k)
	for i := 0; i < k; i++ {
		j := len(que) - 1
		for j >= 0 && nums[que[j]] < nums[i] {
			j--
		}
		que = append(que[:j+1], i)
	}
	ans := make([]int, 0, len(nums)-k+1)
	ans = append(ans, nums[que[0]])
	for i := k; i < len(nums); i++ {
		if i == que[0] {
			que = que[1:]
		}
		j := len(que) - 1
		for j >= 0 && nums[que[j]] < nums[i] {
			j--
		}
		que = append(que[:j+1], nums[i])
		ans = append(ans, que[0])
	}
	return ans
}
