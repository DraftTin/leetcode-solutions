package main

type MyCalendar struct {
	events [][2]int
}

func Constructor() MyCalendar {
	events := [][2]int{}
	return MyCalendar{events: events}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, event := range this.events {
		if start < event[1] && start >= event[0] || end < event[1] && end > event[0] || start <= event[0] && end >= event[1] {
			return false
		}
	}
	this.events = append(this.events, [2]int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
