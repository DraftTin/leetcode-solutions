package main

func checkValidString(s string) bool {
	count := 0
	return helperSolution(count, s, 0)
}

func helperSolution(count int, s string, cur int) bool {
	if cur == len(s) {
		return count == 0
	}
	if count == 0 {
		if s[cur] == ')' {
			return false
		} else if s[cur] == '(' {
			count++
			return helperSolution(count, s, cur+1)
		} else {
			if helperSolution(count+1, s, cur+1) == true {
				return true
			}
			return helperSolution(count, s, cur+1)
		}
	}
	if s[cur] == ')' {
		return helperSolution(count-1, s, cur+1)
	} else if s[cur] == '(' {
		return helperSolution(count+1, s, cur+1)
	} else {
		if helperSolution(count+1, s, cur+1) == true {
			return true
		}
		if helperSolution(count-1, s, cur+1) == true {
			return true
		}
		return helperSolution(count, s, cur+1)
	}
}
