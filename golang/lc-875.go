package main

import "math"

func minEatingSpeed(piles []int, h int) int {
	maxVal := 0
	for _, pile := range piles {
		maxVal = max(maxVal, pile)
	}
	l, r := 1, maxVal
	res := math.MaxInt
	for l <= r {
		mid := (l + r) / 2
		cost := 0
		for _, pile := range piles {
			cost += pile / mid
			if pile%mid != 0 {
				cost += 1
			}
			if cost > h {
				break
			}
		}
		if cost > h {
			l = mid + 1
		} else {
			r = mid - 1
			res = min(res, mid)
		}
	}
	return res
}
