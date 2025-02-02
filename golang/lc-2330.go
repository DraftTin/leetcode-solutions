package main

func makePalindrome(s string) bool {
	half := len(s) / 2
	count := 0
	for i := 0; i <= half; i++ {
		if s[i] != s[len(s)-1-i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
}
