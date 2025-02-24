package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mp := map[byte]string{}
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"
	curStr := make([]byte, len(digits))
	ans := []string{}
	dfs17(digits, 0, curStr, &ans, mp)
	return ans
}

func dfs17(digits string, cur int, curStr []byte, ans *[]string, mp map[byte]string) {
	if cur == len(digits) {
		*ans = append(*ans, string(curStr))
		return
	}
	s := mp[digits[cur]]
	for i := 0; i < len(s); i++ {
		curStr[cur] = s[i]
		dfs17(digits, cur+1, curStr, ans, mp)
	}
}
