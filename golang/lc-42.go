package main

func trap(height []int) int {
	sta := []int{}
	i := 0
	res := 0
	for i < len(height) {
		for len(sta) != 0 && height[sta[len(sta)-1]] < height[i] {
			midHeight := height[sta[len(sta)-1]]
			sta = sta[:len(sta)-1]
			if len(sta) != 0 {
				l := sta[len(sta)-1]
				leftHeight := height[l]
				tmpHeight := min(leftHeight, height[i])
				res += (tmpHeight - midHeight) * (i - l - 1)
			}
		}
		sta = append(sta, i)
		i++
	}
	return res
}
