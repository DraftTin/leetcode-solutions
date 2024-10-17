package main

func minAddToMakeValid(s string) int {
	count := 0
	res := 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if count == 0 {
			res++
		} else {
			count--
		}
	}
	return res + count
}
