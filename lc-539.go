package main

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	timeList := make([]int, 0, len(timePoints))
	for _, timePoint := range timePoints {
		mins, _ := strconv.Atoi(timePoint[:2])
		hours, _ := strconv.Atoi(timePoint[3:])
		timeList = append(timeList, mins+hours*60)
	}
	sort.Ints(timeList)
	res := 0
	for i := 0; i < len(timeList)-1; i++ {
		res = min(res, int(math.Abs(float64(timeList[i]-timeList[i+1]))))
	}
	return res
}
