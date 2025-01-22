package main

func findThePrefixCommonArray(A []int, B []int) []int {
	dictA, dictB := make(map[int]bool), make(map[int]bool)
	res := make([]int, 0, len(A))
	count := 0
	for i := 0; i < len(A); i++ {
		dictA[A[i]] = true
		dictB[B[i]] = true
		if A[i] == B[i] {
			count++
		} else {
			if dictA[B[i]] == true {
				count++
			}
			if dictB[A[i]] == true {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}
