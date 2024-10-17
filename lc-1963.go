package main

func minSwaps(s string) int {
	cnt := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			if cnt == 0 {
				res++
				cnt++
			} else {
				cnt--
			}
		} else {
			cnt++
		}
	}
	return res
}
