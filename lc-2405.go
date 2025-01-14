package main

func partitionString(s string) int {
	res := 1
	dict := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == true {
			dict = make(map[byte]bool)
			res++
		}
		dict[s[i]] = true
	}
	return res
}
