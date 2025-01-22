package main

import "strconv"

func compress(chars []byte) int {
	i := 0
	pos := 0
	res := 0
	for i < len(chars) {
		j := i + 1
		for j < len(chars) && chars[j] == chars[i] {
			j++
		}
		res++
		chars[pos] = chars[i]
		pos++
		if j-i > 1 {
			count := strconv.Itoa(j - i)
			res += len(count)
			copy(chars[pos:], []byte(count))
			pos += len(count)
		}
		i = j
	}
	return res
}
