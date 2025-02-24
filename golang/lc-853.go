package main

import "sort"

func carFleet(target int, position []int, speed []int) int {
	sta := [][2]int{}
	cars := make([][2]int, len(position))
	for i := range position {
		cars[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	i := 0
	for i < len(cars) {
		sta = append(sta, cars[i])
		for len(sta) != 1 {
			t1 := float64((target - sta[len(sta)-1][0])) / float64(sta[len(sta)-1][1])
			t2 := float64((target - sta[len(sta)-2][0])) / float64(sta[len(sta)-2][1])
			if t1 <= t2 {
				sta = sta[:len(sta)-1]
			} else {
				break
			}
		}
		i++
	}
	return len(sta)
}
