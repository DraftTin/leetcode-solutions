package main

type TopVotedCandidate struct {
	PersonSeries []int
	TimeSeries   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	personDict := map[int]int{}
	maxVote := 0
	personSeries := make([]int, 0, len(times))
	for i, _ := range times {
		personDict[persons[i]]++
		if personDict[persons[i]] >= maxVote {
			personSeries = append(personSeries, persons[i])
			maxVote = personDict[persons[i]]
		} else {
			personSeries = append(personSeries, personSeries[len(personSeries)-1])
		}
	}
	return TopVotedCandidate{PersonSeries: personSeries, TimeSeries: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.TimeSeries)
	for l <= r {
		mid := (l + r) / 2
		if t >= this.TimeSeries[mid] && (mid == len(this.TimeSeries)-1 || t < this.TimeSeries[mid+1]) {
			return this.PersonSeries[mid]
		} else if t < this.TimeSeries[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
