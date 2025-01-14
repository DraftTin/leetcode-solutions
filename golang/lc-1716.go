func totalMoney(n int) int {
	n1, n2 := n/7, n%7
	ans := (n1 * (n1 - 1)) / 2 * 7
	ans += n1 * 28
	ans += (n2*(n2+1))/2 + n1*n2
	return ans
}
