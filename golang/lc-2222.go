package main

func numberOfWays(s string) (ans int64) {
	countZero := 0
	for i := range s {
		if s[i] == '0' {
			countZero++
		}
	}
	countOne := len(s) - countZero

	numZero, numOne := 0, 0
	for i := range s {
		if s[i] == '0' {
			numZero++
			ans += int64(numOne * (countOne - numOne))
		} else {
			numOne++
			ans += int64(numZero * (countZero - numZero))
		}
	}

	return
}
