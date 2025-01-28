package main

func largestAltitude(gain []int) int {
	maxHeight := 0
	curHeight := 0
	for i := 0; i < len(gain); i++ {
		curHeight += gain[i]
		maxHeight = max(maxHeight, curHeight)
	}
	return maxHeight
}
