package main

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	for _, num := range nums {
		count[num]++
	}
	c := 0
	for i := 0; i < 101; i++ {
		c += count[i]
		count[i] = c - count[i]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = count[nums[i]]
	}
	return res
}
