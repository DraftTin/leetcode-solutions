package main

import "fmt"

func findFirstGreater(nums []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = -1 // 默认设置为-1
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = nums[i] // 当前数是前面第一个比它大的数
			stack = stack[:len(stack)-1]          // 弹出栈顶元素
		}
		stack = append(stack, i) // 将当前索引入栈
	}

	// 处理没有前面比它大的数的情况
	for i := range result {
		if result[i] == -1 { // 没有前面比它大的数
			result[i] = result[0] // 输出第一个比它大的数
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 5, 2, 7, 9}
	result := findFirstGreater(nums)
	fmt.Println(result) // 输出结果: [1 3 1 5 1 7]
}
