package main

import "strconv"

func getHint(secret string, guess string) string {
	secretDict, guessDict := make(map[byte]int), make(map[byte]int)
	bull := 0
	cow := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
		} else {
			secretDict[secret[i]]++
			guessDict[guess[i]]++
			if guessDict[secret[i]] > 0 {
				guessDict[secret[i]]--
				secretDict[secret[i]]--
				cow++
			}
			if secretDict[guess[i]] > 0 {
				secretDict[guess[i]]--
				guessDict[guess[i]]--
				cow++
			}
		}
	}
	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}
