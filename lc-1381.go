package main

type CustomStack struct {
	capacity int
	stack    []int
}

func Constructor(maxSize int) CustomStack {
	stack := make([]int, 0, maxSize)
	customStack := CustomStack{capacity: maxSize, stack: stack}
	return customStack
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) == this.capacity {
		return
	}
	this.stack = append(this.stack, x)
}

func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	res := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	end := min(len(this.stack), k)
	for i := 0; i < end; i++ {
		this.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
