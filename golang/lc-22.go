package main

func generateParenthesis(n int) []string {
	mp := make(map[int][]string)
	mp[2] = []string{"()"}
	mp[0] = []string{}
	curString := make([]uint8, 2*n)
	for i := 4; i <= 2*n; i += 2 {
		for j := 1; j < i; j += 2 {
			curString[0] = '('
			curString[j] = ')'
			for _, substr := range mp[j-1] {
				helper(1, j, curString, substr)
				for _, substr := range mp[i-j-1] {
					helper(j+1, i-1, curString, substr)
					mp[i] = append(mp[i], string(curString))
				}
			}
		}
	}
	return mp[n*2]
}

func helper(i, j int, s []uint8, substr string) {
	for k := i; i <= j; k++ {
		s[k] = substr[k-i]
	}
}
