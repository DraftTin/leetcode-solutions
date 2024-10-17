package main

func minimumSteps(s string) int64 {
	zeroNum := 0
	res := int64(0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroNum++
		} else {
			res += int64(zeroNum)
		}
	}
	return res
}
