package main

type Stack []int

type MyQueue struct {
	sta1, sta2 Stack
}

func (this *Stack) Pop() int {
	old := *this
	item := old[len(old)-1]
	*this = old[0 : len(old)-1]
	return item
}

func (this *Stack) Push(val int) {
	*this = append(*this, val)
}

func Constructor() MyQueue {
	return MyQueue{}

}

func (this *MyQueue) Push(x int) {
	for len(this.sta1) != 0 {
		val := this.sta1.Pop()
		this.sta2.Push(val)
	}
	this.sta1.Push(x)
	for len(this.sta2) != 0 {
		val := this.sta2.Pop()
		this.sta1.Push(val)
	}
}

func (this *MyQueue) Pop() int {
	return this.sta1.Pop()
}

func (this *MyQueue) Peek() int {
	return this.sta1[len(this.sta1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.sta1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
