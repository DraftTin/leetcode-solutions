package main

import "strconv"

// String Manipulation + Digit Operatio
func getLucky(s string, k int) int {
	digits := ""
	for i := 0; i < len(s); i++ {
		digits += strconv.Itoa(int(s[i] - 'a'))
	}

	for i := 0; i < k; i++ {
		val := 0
		for _, digit := range digits {
			val += int(digit - '0')
		}
		digits = strconv.Itoa(val)
	}
	res, _ := strconv.Atoi(digits)
	return res
}
