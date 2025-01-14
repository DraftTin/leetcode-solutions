package main

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		if i*(1+i) == (n-i+1)*(i+n) {
			return i
		}
	}
	return -1
}
