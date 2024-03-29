package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := 0
		if s[i] != '0' {
			tmp = cur
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				tmp += pre
			}
		} else {
			if s[i-1] == '1' || s[i-1] == '2' {
				tmp += pre
			}
		}
		pre, cur = cur, tmp
	}
	return cur
}
