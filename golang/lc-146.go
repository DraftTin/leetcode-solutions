package main

import "fmt"

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	LRU      *Node
	MRU      *Node
	Cache    map[int]*Node
	Capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.Cache = make(map[int]*Node)
	cache.Capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	if this.Cache[key] != nil {
		this.PutHead(this.Cache[key])
		return this.Cache[key].Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Cache[key] != nil {
		this.Cache[key].Val = value
	} else {
		if len(this.Cache) == this.Capacity {
			this.Pop()
		}
		this.Cache[key] = &Node{Key: key, Val: value}
	}
	this.PutHead(this.Cache[key])
}

func (this *LRUCache) PutHead(node *Node) {
	if this.MRU == nil {
		this.MRU = node
		this.LRU = node
		return
	}
	if this.MRU == node {
		return
	}
	if this.LRU == node {
		this.LRU = this.LRU.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = this.MRU
	this.MRU.Prev = node
	this.MRU = node
}

func (this *LRUCache) Pop() {
	fmt.Println(this.Capacity, len(this.Cache), this.LRU)
	delete(this.Cache, this.LRU.Key)
	if this.LRU == this.MRU {
		this.LRU = nil
		this.MRU = nil
		return
	}
	this.LRU = this.LRU.Prev
	this.LRU.Next = nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
