package main

func maxFrequencyElements(nums []int) int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}
	maxCount := 0
	total := 0
	for _, val := range frequencies {
		if maxCount < val {
			total = val
			maxCount = val
		} else if maxCount == val {
			total += val
		}
	}
	return total
}
