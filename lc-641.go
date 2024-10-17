package main

type MyCircularDeque struct {
	queue    []int
	front    int
	rear     int
	capacity int
	size     int
}

// circular double-ended queue
func Constructor(k int) MyCircularDeque {
	que := make([]int, k)
	circularDeque := MyCircularDeque{queue: que, capacity: k, front: 0, rear: k - 1}
	return circularDeque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size >= this.capacity {
		return false
	}
	this.front = (this.front - 1 + this.capacity) % this.capacity
	this.queue[this.front] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size >= this.capacity {
		return false
	}
	this.rear = (this.rear + 1) % this.capacity
	this.queue[this.rear] = value
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 {
		return false
	}
	this.front = (this.front + 1) % this.capacity
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}
	this.rear = (this.rear - 1 + this.capacity) % this.capacity
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}
	return this.queue[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}
	return this.queue[this.rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
