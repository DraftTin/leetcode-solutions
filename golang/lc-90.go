package main

import (
	"sort"
	"strconv"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	set := map[string]bool{"": true}
	for i := 0; i < len(nums); i++ {
		tmpSet := map[string]bool{}
		for key, _ := range set {
			numArr := strings.Split(key, ",")
			numArr = append(numArr, strconv.Itoa(nums[i]))
			sort.Strings(numArr)
			newKey := strings.Join(numArr, ",")
			tmpSet[newKey] = true
		}
		for key, _ := range tmpSet {
			set[key] = true
		}
	}
	ans := make([][]int, 0, len(set)-1)
	for key, _ := range set {
		numArr := strings.Split(key, ",")
		item := make([]int, len(numArr)-1)
		for i := 0; i < len(item); i++ {
			item[i], _ = strconv.Atoi(numArr[i+1])
		}
		ans = append(ans, item)
	}
	return ans
}
