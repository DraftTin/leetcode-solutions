package main

func longestConsecutive(nums []int) int {
	numDict := map[int]int{}
	maxVal := 0
	father := map[int]int{}
	for i := 0; i < len(nums); i++ {
		father[i] = i
	}
	numDict[-1] = 0
	for _, num := range nums {
		if numDict[num] != 0 {
			continue
		}
		f1, f2 := findFather(father, num-1), findFather(father, num+1)
		numDict[num] = 1
		if f1 != -1 {
			numDict[num] += numDict[f1]
			father[f1] = num
		}
		if f2 != -1 {
			numDict[num] += numDict[f2]
			father[f2] = num
		}
		maxVal = max(maxVal, numDict[num])
	}
	return maxVal
}

func findFather(father map[int]int, pos int) int {
	if _, ok := father[pos]; ok == false {
		return -1
	}
	for father[pos] != pos {
		father[pos] = findFather(father, father[pos])
	}
	return father[pos]
}
