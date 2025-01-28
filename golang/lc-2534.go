package main

func timeTaken(arrival []int, state []int) []int {
	enterQ, exitQ := []int{}, []int{}
	t := 0
	i := 0
	prev := -1
	res := make([]int, len(arrival))
	for true {
		for i < len(arrival) && arrival[i] == t {
			switch state[i] {
			case 0:
				enterQ = append(enterQ, i)
			case 1:
				exitQ = append(exitQ, i)
			}
			i++
		}
		tmp := -1
		switch {
		case len(enterQ) > 0 && len(exitQ) > 0:
			if prev == -1 || prev == 1 {
				tmp = exitQ[0]
				exitQ = exitQ[1:]
				prev = 1
			} else {
				tmp = enterQ[0]
				enterQ = enterQ[1:]
				prev = 0
			}
		case len(enterQ) > 0:
			tmp = enterQ[0]
			enterQ = enterQ[1:]
			prev = 0
		case len(exitQ) > 0:
			tmp = exitQ[0]
			exitQ = exitQ[1:]
			prev = 1
		default:
			prev = -1
		}
		if tmp != -1 {
			res[tmp] = t
		}
		if i == len(arrival) && len(exitQ) == 0 && len(enterQ) == 0 {
			break
		}

		t++
	}
	return res
}
