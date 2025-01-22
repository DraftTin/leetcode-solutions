package main

func nextGreaterElements(nums []int) []int {
	stack := [][2]int{}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	set := make(map[int]bool)
	for i := 0; i < 2; i++ {
		for i, num := range nums {
			for len(stack) != 0 && stack[len(stack)-1][0] < num {
				res[stack[len(stack)-1][1]] = num
				stack = stack[:len(stack)-1]
			}
			if set[i] == false {
				set[i] = true
				stack = append(stack, [2]int{nums[i], i})
			}
		}
	}
	return res
}
