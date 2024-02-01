package main

func dailyTemperatures(temperatures []int) []int {
	sta := []int{}
	ans := make([]int, 0, len(temperatures))
	index := 0
	for i, temperature := range temperatures {
		for len(sta) != 0 && temperatures[sta[len(sta)-1]] < temperature {
			ans[index] = i - sta[len(sta)-1]
			index++
			sta = sta[:len(sta)-1]
		}
		sta = append(sta, i)
	}
	return ans
}
