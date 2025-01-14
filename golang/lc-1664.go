func waysToMakeFair(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	oddSum, evenSum := make([]int, n), make([]int, n)
	sum1, sum2 := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum2 += nums[i]
			evenSum[i] = sum2
		} else {
			sum1 += nums[i]
			oddSum[i] = sum1
		}
	}
	ans := 0
	totalSum1, totalSum2 := 0, 0
	if n%2 == 1 {
		totalSum2 = evenSum[n-1]
		totalSum1 = oddSum[n-2]
	} else {
		totalSum2 = evenSum[n-2]
		totalSum1 = oddSum[n-1]
	}
	for i := 0; i < n; i++ {
		sum1, sum2 = 0, 0
		if i%2 == 0 {
			if i != 0 {
				sum2 += evenSum[i-2] - oddSum[i-1]
				sum1 += oddSum[i-1]
			}
			sum1 += totalSum2 - evenSum[i]
			sum2 += totalSum1
		} else {
			sum2 += evenSum[i-1] + totalSum1 - oddSum[i]
			sum1 += totalSum2 - evenSum[i-1]
			if i != 1 {
				sum1 += oddSum[i-2]
			}
		}
		if sum1 == sum2 {
			ans++
		}
	}
	return ans
}
