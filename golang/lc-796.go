package main

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[:i] == goal[len(goal)-i:] && s[i:] == goal[:len(goal)-i] {
			return true
		}
	}
	return false

}
