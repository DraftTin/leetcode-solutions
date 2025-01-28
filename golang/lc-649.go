package main

func predictPartyVictory(senate string) string {
	banned := make([]bool, len(senate))
	i := 0
	dict := map[byte]string{'R': "Radiant", 'D': "Dire"}
	for true {
		for i < len(senate) && banned[i] == true {
			i++
		}
		if i == len(senate) {
			i = 0
			continue
		}
		j := i + 1
		for j < len(senate) && (banned[j] || senate[j] == senate[i]) {
			j++
		}
		if j == len(senate) {
			j = 0
			for j < i && (banned[j] || senate[j] == senate[i]) {
				j++
			}
		}
		if j == i {
			return dict[senate[i]]
		} else {
			banned[j] = true
		}
		i++
	}
	return ""
}
