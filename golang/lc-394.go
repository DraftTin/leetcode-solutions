package main

import (
	"strconv"
)

// recursion
func decodeString(s string) string {
	return helper394(s)
}

func helper394(s string) string {
	posDict := make(map[int]int)
	sta := []int{}
	for i, c := range s {
		if c == '[' {
			sta = append(sta, i)
		} else if c == ']' {
			top := sta[len(sta)-1]
			sta = sta[:len(sta)-1]
			posDict[top] = i
		}
	}
	res := ""
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			num := []byte{}
			k := i
			for s[k] >= '0' && s[k] <= '9' {
				num = append(num, s[k])
				k++
			}
			val, _ := strconv.Atoi(string(num))
			subString := helper394(s[k+1 : posDict[k]])
			for val > 0 {
				res += subString
				val--
			}
			i = posDict[k] + 1
		} else {
			res += string(s[i])
			i++
		}
	}
	return res
}
