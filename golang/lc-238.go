package main

func productExceptSelf(nums []int) []int {
	leftProducts, rightProducts := make([]int, len(nums)), make([]int, len(nums))
	val := 1
	for i := 0; i < len(nums); i++ {
		val *= nums[i]
		leftProducts[i] = val
	}
	val = 1
	for i := len(nums) - 1; i >= 0; i-- {
		val *= nums[i]
		rightProducts[i] = val
	}

	res := make([]int, len(nums))
	res[0] = rightProducts[1]
	res[len(res)-1] = leftProducts[len(res)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = rightProducts[i+1] * leftProducts[i-1]
	}
	return res
}
