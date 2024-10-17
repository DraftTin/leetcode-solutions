package main

// handle the situation where the remainder == (k - remainder) % k
func canArrange(arr []int, k int) bool {
	remainderFreq := make([]int, k)

	for _, num := range arr {
		remainder := ((num % k) + k) % k
		remainderFreq[remainder]++
	}

	for i := 0; i <= k/2; i++ {
		if i == (k-i)%k {
			if remainderFreq[i]%2 != 0 {
				return false
			}
		} else {
			if remainderFreq[i] != remainderFreq[k-i] {
				return false
			}
		}
	}

	return true
}
