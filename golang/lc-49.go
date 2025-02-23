package main

func groupAnagrams(strs []string) [][]string {
	resSet := make(map[[26]byte][]int)
	countList := make([][26]byte, len(strs))
	for i, str := range strs {
		for _, c := range str {
			countList[i][c-'a']++
		}
		resSet[countList[i]] = append(resSet[countList[i]], i)
	}
	res := [][]string{}
	for _, set := range resSet {
		tmp := []string{}
		for _, ind := range set {
			tmp = append(tmp, strs[ind])
		}
		res = append(res, tmp)
	}
	return res
}
