package main

func checkZeroOnes(s string) bool {
	maxOne, maxZero := 0, 0
	i := 0
	for i < len(s) {
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if s[i] == '1' {
			maxOne = max(maxOne, j-i)
		} else {
			maxZero = max(maxZero, j-i)
		}
		i = j
	}
	return maxOne > maxZero
}
