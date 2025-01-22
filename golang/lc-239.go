package main

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	que := []int{}
	for i := 0; i < k; i++ {
		j := len(que) - 1
		for j >= 0 && nums[que[j]] < nums[i] {
			j--
		}
		que = append(que[:j+1], i)
	}
	ans = append(ans, nums[que[0]])
	for i := k; i < len(nums); i++ {
		if i-k == que[0] {
			que = que[1:]
		}
		j := len(que) - 1
		for j >= 0 && nums[que[j]] < nums[i] {
			j--
		}
		que = append(que[:j+1], i)
		ans = append(ans, nums[que[0]])
	}
	return ans
}
