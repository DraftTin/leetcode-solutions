package main

import "strconv"

// Greedy Algorithm
func maximumSwap(num int) int {
	res := []byte(strconv.Itoa(num))
	maxVal := res[len(res)-1]
	index := len(res) - 1
	ind1, ind2 := -1, -1
	for i := len(res) - 2; i >= 0; i-- {
		if maxVal > res[i] {
			ind1 = i
			ind2 = index
		} else if maxVal < res[i] {
			maxVal = res[i]
			index = i
		}
	}
	resVal := num
	if ind1 != -1 {
		res[ind1], res[ind2] = res[ind2], res[ind1]
		resVal, _ = strconv.Atoi(string(res))
	}
	return resVal
}
