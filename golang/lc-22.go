package main

func generateParenthesis(n int) []string {
	mp := map[int][]string{}
	mp[0] = []string{""}
	mp[2] = []string{"()"}
	curStr := make([]byte, 2*n)
	for i := 3; i < n*2; i += 2 {
		for j := 1; j <= i; j += 2 {
			curStr[0], curStr[j] = '(', ')'
			for _, substr := range mp[j-1] {
				copy(curStr[1:j+1], []byte(substr))
				for _, substr := range mp[i-j] {
					copy(curStr[j+1:i+1], []byte(substr))
					mp[i+1] = append(mp[i+1], string(curStr[:i+1]))
				}
			}
		}
	}
	return mp[2*n]
}
