package main

type RecentCounter struct {
	Ind     int
	PingArr []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.PingArr = append(this.PingArr, t)
	for this.PingArr[this.Ind] < t-3000 {
		this.Ind++
	}
	return len(this.PingArr) - this.Ind
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
