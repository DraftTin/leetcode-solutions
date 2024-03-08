package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	ans := [][]int{}
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][1] < secondList[j][0] || secondList[j][1] < firstList[i][0] {
			for i < len(firstList) && firstList[i][1] < secondList[j][0] {
				i++
			}
			if i == len(firstList) {
				break
			}
			for j < len(secondList) && secondList[j][1] < firstList[i][0] {
				j++
			}
			if j == len(secondList) {
				break
			}
			continue
		}
		var l, r int
		if firstList[i][0] < secondList[j][0] {
			l = secondList[j][0]
		} else {
			l = firstList[i][0]
		}
		if firstList[i][1] < secondList[j][1] {
			r = firstList[i][1]
			i++
		} else if firstList[i][1] > secondList[j][1] {
			r = secondList[j][1]
			j++
		} else {
			r = secondList[j][1]
			i++
			j++
		}
		ans = append(ans, []int{l, r})
	}
	return ans
}
