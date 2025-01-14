func subsets(nums []int) [][]int {
	ans := [][]int{[]int{}}
	for _, num := range nums {
		tmp := ans
		for i, _ := range ans {
			item := make([]int, len(ans[i])+1)
			copy(item, ans[i])
			item[len(item)-1] = num
			tmp = append(tmp, item)
		}
		ans = tmp
	}
	return ans
}
