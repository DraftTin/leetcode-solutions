package main

import "math"

// binary search ***
func repairCars(ranks []int, cars int) int64 {
	left, right := int64(0), int64(ranks[0]*cars*cars)
	for left <= right {
		mid := (left + right) / 2
		if solve(mid, ranks, cars) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func solve(time int64, ranks []int, cars int) bool {
	for _, rank := range ranks {
		tmp := int(math.Sqrt(float64(time / int64(rank))))
		cars -= tmp
		if cars <= 0 {
			return true
		}
	}
	return false
}
