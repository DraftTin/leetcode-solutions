package main

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return gcdOfStrings(str2, str1)
	}
	i := 0
	for i+len(str2) <= len(str1) {
		if str1[i:len(str2)+i] != str2 {
			return ""
		}
		i += len(str2)
	}
	if i == len(str1) {
		return str2
	}
	return gcdOfStrings(str2, str1[i:])
}
