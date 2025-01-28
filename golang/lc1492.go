package main

func kthFactor(n int, k int) int {
	factors := make([]int, 0, n)
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if n/i+count < k {
				return -1
			}
			factors = append(factors, i)
			count++
			if count == k {
				return factors[len(factors)-1]
			}
		}
	}
	return -1
}
