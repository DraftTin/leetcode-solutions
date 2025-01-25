package main

import "strconv"

func countAndSay(n int) string {
	res := []byte{'1'}
	for i := 1; i < n; i++ {
		j := 0
		newRes := make([]byte, 0, len(res))
		for j < len(res) {
			k := j + 1
			for k < len(res) && res[k] == res[j] {
				k++
			}
			newRes = append(newRes, []byte(strconv.Itoa(k-j))...)
			newRes = append(newRes, res[j])
			j = k
		}
		res = newRes
	}
	return string(res)
}
