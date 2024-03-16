package main

func findMaxLength(nums []int) int {
	mp := make(map[int]int)
	ans := 0
	sum := 0
	for i, num := range nums {
		if num == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		if sum == 0 {
			ans = i + 1
		} else if pos, ok := mp[sum]; ok == true {
			ans = max(ans, i-pos)
		} else {
			mp[sum] = i
		}
	}
	return ans
}
