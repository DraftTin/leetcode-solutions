package main

// DP
func countSubstrings(s string) int {
	oddMarks := make([]bool, len(s))
	evenMarks := make([]bool, len(s))
	res := len(s)
	for i := 0; i < len(s)-1; i++ {
		oddMarks[i] = true
		if s[i] == s[i+1] {
			evenMarks[i] = true
			res++
		}
	}
	oddMarks[len(s)-1] = true
	for i := 3; i <= len(s); i++ {
		for j := i - 1; j < len(s); j++ {
			if s[j-i+1] != s[j] {
				if i%2 == 1 {
					oddMarks[j-i/2] = false
				} else {
					evenMarks[j-i/2] = false
				}
			} else {
				if i%2 == 1 && oddMarks[j-i/2] == true {
					res++
				}
				if i%2 == 0 && evenMarks[j-i/2] == true {
					res++
				}
			}
		}
	}
	return res
}
