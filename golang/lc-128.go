package main

func longestConsecutive(nums []int) int {
	mp := map[int]int{}
	n := len(nums)
	fathers := make([]int, n)
	counter := make([]int, n)
	for i := 0; i < n; i++ {
		fathers[i] = i
		counter[i] = 1
	}
	ans := 0
	for i, num := range nums {
		if mp[num] == 0 {
			mp[num] = i + 1
			if mp[num-1] != 0 {
				union(fathers, counter, mp[num]-1, mp[num-1]-1)
			}
			if mp[num+1] != 0 {
				union(fathers, counter, mp[num]-1, mp[num+1]-1)
			}
			tmp := getCount(fathers, counter, i)
			ans = max(ans, tmp)
		}
	}
	return ans
}

func union(fathers []int, counter []int, i, j int) {
	pi := findFather(fathers, i)
	pj := findFather(fathers, j)
	fathers[pj] = pi
	counter[pi] = counter[pi] + counter[pj]
}

func getCount(fathers []int, counter []int, i int) int {
	p := findFather(fathers, i)
	return counter[p]
}

func findFather(fathers []int, i int) int {
	tmp := i
	for fathers[i] != i {
		i = fathers[i]
	}
	fathers[tmp] = i
	return i
}
