package main

// Greedy Algorithm: match as long as possible
func checkPalindromeFormation(a string, b string) bool {
	i := 0
	for i <= len(a)-1-i {
		if a[i] != b[len(b)-1-i] {
			break
		}
		i++
	}
	if i >= len(a)-1-i {
		return true
	} else {
		j, k := i, len(b)-1-i
		for j < k {
			if b[j] != b[k] {
				break
			}
			j++
			k--
		}
		if j >= k {
			return true
		}
		j, k = i, len(b)-1-i
		for j < k {
			if a[j] != a[k] {
				break
			}
			j++
			k--
		}
		if j >= k {
			return true
		}
	}
	i = 0
	for i < len(a) {
		if b[i] != a[len(b)-1-i] {
			break
		}
		i++
	}
	if i >= len(a)-1-i {
		return true
	} else {
		j, k := i, len(b)-1-i
		for j < k {
			if a[j] != a[k] {
				break
			}
			j++
			k--
		}
		if j >= k {
			return true
		}
		j, k = i, len(b)-1-i
		for j < k {
			if b[j] != b[k] {
				break
			}
			j++
			k--
		}
		if j >= k {
			return true
		}
	}
	return false
}
