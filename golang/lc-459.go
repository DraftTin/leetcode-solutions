package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	newS := s[1:] + s[:len(s)-1]
	index := strings.Index(newS, s)
	if index == -1 {
		return false
	}
	return true
}
