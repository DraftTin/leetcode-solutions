package main

func longestConsecutive(nums []int) int {
	parents := make(map[int]int)
	count := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := count[num]; ok {
			continue
		}
		parents[num] = num
		count[num] = 1
		if _, ok := count[num-1]; ok {
			union(parents, num, num-1, count)
		}
		if _, ok := count[num+1]; ok {
			union(parents, num, num+1, count)
		}
		p := findParent(parents, num)
		res = max(res, count[p])
	}
	return res
}

func union(parents map[int]int, lval, rval int, count map[int]int) {
	p1, p2 := findParent(parents, lval), findParent(parents, rval)
	if p1 == p2 {
		return
	}
	if count[p1] > count[p2] {
		parents[p2] = p1
		count[p1] += count[p2]
	} else {
		parents[p1] = p2
		count[p2] += count[p1]
	}
}

func findParent(parents map[int]int, val int) int {
	if val != parents[val] {
		parents[val] = findParent(parents, parents[val])
	}
	return parents[val]
}
