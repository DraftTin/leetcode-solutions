package main

func reverseWords(s string) string {
	ans := []string{}
	cur := 0
	for s[cur] == ' ' {
		cur++
	}
	for cur < len(s) {
		start := cur
		for cur < len(s) && s[cur] != ' ' {
			cur++
		}
		ans = append(ans, s[start:cur])
		for cur < len(s) && s[cur] == ' ' {
			cur++
		}
	}
	reversedStr := ""
	for i := len(ans) - 1; i >= 0; i-- {
		reversedStr = reversedStr + ans[i] + " "
	}
	return reversedStr[:len(reversedStr)-1]
}
