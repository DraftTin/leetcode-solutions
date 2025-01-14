package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsStr := make([]string, 0, len(nums))
	for _, num := range nums {
		numsStr = append(numsStr, strconv.Itoa(num))
	}
	sort.Slice(numsStr, func(i, j int) bool {
		return numsStr[i]+numsStr[j] >= numsStr[j]+numsStr[i]
	})
	return strings.Join(numsStr, "")
}
