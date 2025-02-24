package main

func dailyTemperatures(temperatures []int) []int {
	sta := []int{}
	res := make([]int, len(temperatures))
	i := 0
	for i < len(temperatures) {
		for len(sta) != 0 && temperatures[i] > temperatures[sta[len(sta)-1]] {
			ind := sta[len(sta)-1]
			res[ind] = i - ind
			sta = sta[:len(sta)-1]
		}
		sta = append(sta, i)
		i++
	}
	return res
}
