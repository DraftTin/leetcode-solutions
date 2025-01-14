package main

// Bit Manipulation + backtrack
func findKthBit(n int, k int) byte {
	if k == 1 || n <= 1 {
		return '0'
	}
	bitsNum := 1
	for i := 0; i < n-1; i++ {
		bitsNum = 2*bitsNum + 1
	}
	if k-1 == bitsNum/2 {
		return '1'
	} else if k-1 > bitsNum/2 {
		tmp := ((findKthBit(n-1, bitsNum-k+1)+1)%'0'+'0')%2 + '0'
		return tmp
	}
	return findKthBit(n-1, k)
}
