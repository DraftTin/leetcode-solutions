package main

func longestPalindrome(s string) string {
	maxLen := 0
	left := -1
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxLen {
			maxLen = r - l - 1
			left = l + 1
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxLen {
			maxLen = r - l - 1
			left = l + 1
		}
	}
	return s[left : left+maxLen]
}
