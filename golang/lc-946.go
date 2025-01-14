package main

func validateStackSequences(pushed []int, popped []int) bool {
	numMap := make(map[int]int)
	for i := 0; i < len(pushed); i++ {
		numMap[pushed[i]] = i
		pushed[i] = i
	}
	for i := 0; i < len(popped); i++ {
		popped[i] = numMap[popped[i]]
	}
	topList := []int{}
	poppedSet := map[int]bool{}
	for i := 0; i < len(popped); i++ {
		poppedSet[popped[i]] = true
		if len(topList) == 0 {
			topList = append(topList, popped[i]-1)
		} else if topList[len(topList)-1] == popped[i] {
			if poppedSet[popped[i]-1] == true {
				topList = topList[:len(topList)-1]
			} else {
				topList[len(topList)-1]--
			}
		} else if topList[len(topList)-1] < popped[i] {
			if poppedSet[popped[i]-1] == false {
				topList = append(topList, popped[i]-1)
			}
		} else if topList[len(topList)-1] > popped[i] {
			return false
		}
	}
	return true
}
