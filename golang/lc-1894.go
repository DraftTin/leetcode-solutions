package main

// Solutio: prefix sum + simulation
func chalkReplacer(chalk []int, k int) int {
	sumTable := make([]int, len(chalk))
	curSum := 0
	for i := 0; i < len(chalk); i++ {
		curSum += chalk[i]
		sumTable[i] = curSum
	}
	totalSum := sumTable[len(sumTable)-1]
	for i := 0; i < len(sumTable); i++ {
		tmp := k % totalSum
		if tmp < sumTable[i] {
			return i
		}
	}
	return 0
}
