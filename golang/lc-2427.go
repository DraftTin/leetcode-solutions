package main

func commonFactors(a int, b int) int {
	cd := gcd(a, b)
	res := 1
	if cd != 1 {
		res++
	}
	for i := 2; i <= cd/2; i++ {
		if cd%i == 0 {
			res++
		}
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	if a < b {
		return gcd(b, a)
	}
	return gcd(b, a%b)
}
