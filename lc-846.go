package main

import (
	"math"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	numCount := make(map[int]int)
	for _, num := range hand {
		numCount[num]++
	}
	n := len(hand)
	for n != 0 {
		count, flag := helper(numCount, groupSize)
		if !flag {
			return false
		}
		n -= count
	}
	return true
}

func helper(numCount map[int]int, groupSize int) (int, bool) {
	minNum := math.MaxInt32
	for key, count := range numCount {
		if key < minNum && count > 0 {
			minNum = key
		}
	}
	count := numCount[minNum]
	for i := 0; i < groupSize; i++ {
		numCount[i+minNum] -= count
		if numCount[i+minNum] < 0 {
			return 0, false
		}
	}
	return count * groupSize, true
}
