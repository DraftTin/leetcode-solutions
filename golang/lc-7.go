package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	maxInt, minInt := strconv.Itoa(math.MaxInt32), strconv.Itoa(math.MinInt32)
	curVal := strconv.Itoa(x)
	if x < 0 {
		curVal = curVal[1:]
	}
	reversedVal := make([]byte, len(curVal))
	for i := 0; i < len(reversedVal); i++ {
		reversedVal[i] = curVal[len(curVal)-1-i]
	}
	ans := string(reversedVal)
	if len(ans) == len(maxInt) {
		if x > 0 && ans > maxInt || x < 0 && ans > minInt[1:] {
			return 0
		}
	}
	res, _ := strconv.Atoi(ans)
	if x > 0 {
		return res
	}
	return -res
}
