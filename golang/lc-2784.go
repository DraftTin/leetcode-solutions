package main

func isGood(nums []int) bool {
	visited := make(map[int]bool)
	flag := false
	for _, num := range nums {
		if num > len(nums)-1 {
			return false
		}
		if visited[num] == false {
			visited[num] = true
		} else if num != len(nums)-1 || flag == true && num == len(nums)-1 {
			return false
		} else {
			flag = true
		}
	}
	return true

}
