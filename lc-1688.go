package main

func numberOfMatches(n int) int {
	ans := 0
	for n != 1 {
		ans += n / 2
		n -= n / 2
	}
	return ans
}
