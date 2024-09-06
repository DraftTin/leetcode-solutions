package main

// Mathematical Deduction + Greedy Algorithm
func missingRolls(rolls []int, mean int, n int) []int {
	sumVal := (n + len(rolls)) * mean
	mSum := 0
	for i := 0; i < len(rolls); i++ {
		mSum += rolls[i]
	}
	nSum := sumVal - mSum
	res := make([]int, n)
	minCapacity, maxCapacity := n, 6*n
	if nSum < minCapacity || nSum > maxCapacity {
		return []int{}
	}
	rest := nSum - minCapacity
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 0; i < n; i++ {
		if rest >= 5 {
			res[i] = 6
			rest -= 5
		} else {
			res[i] = 1 + rest
			rest = 0
			break
		}
	}
	return res
}
