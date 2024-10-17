package main

func minBitFlips(start int, goal int) int {
	res := 0
	for start != 0 || goal != 0 {
		if start%2 != goal%2 {
			res++
		}
		start /= 2
		goal /= 2
	}
	return res
}
