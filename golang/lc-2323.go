package main

import "sort"

func minimumTime(jobs []int, workers []int) int {
	sort.Ints(jobs)
	sort.Ints(workers)
	res := 0
	for i := 0; i < len(workers); i++ {
		tmp := jobs[i] / workers[i]
		if jobs[i]%workers[i] != 0 {
			tmp++
		}
		res = max(res, tmp)
	}
	return res
}
