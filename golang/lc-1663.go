func getSmallestString(n int, k int) string {
	ans := make([]byte, n)
	rest := k - n + 1
	i := n - 1
	for i >= 0 {
		if rest >= 26 {
			ans[i] = 'z'
			rest -= 25
		} else {
			ans[i] = byte(int('a') + rest - 1)
			rest = 1
		}
		i--
		if rest == 1 {
			break
		}
	}
	for i >= 0 {
		ans[i] = 'a'
		i--
	}
	return string(ans)
}
