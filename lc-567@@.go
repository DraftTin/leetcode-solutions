package main

func checkInclusion(s1 string, s2 string) bool {
	strMap, windowMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		strMap[s1[i]]++
	}
	start := 0
	end := 0
	for end < len(s2) {
		windowMap[s2[end]]++
		for start <= end && strMap[s2[end]] < windowMap[s2[end]] {
			windowMap[s2[start]]--
			start++
		}
		if end-start+1 == len(s1) {
			return true
		}
		end++
	}
	return false
}
