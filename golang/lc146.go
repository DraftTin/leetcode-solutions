// package main
//
// type LRUNode struct {
// 	Prev *LRUNode
// 	Next *LRUNode
// 	Val  int
// 	Key  int
// }
//
// type LRUCache struct {
// 	Cache    map[int]*LRUNode
// 	MRU      *LRUNode
// 	LRU      *LRUNode
// 	Capacity int
// }
//
// func Constructor(capacity int) LRUCache {
// 	this := LRUCache{}
// 	this.Cache = map[int]*LRUNode{}
// 	this.Capacity = capacity
// 	return this
// }
//
// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.Cache[key]; ok {
// 		this.PutHead(node)
// 		return node.Val
// 	}
// 	return -1
// }
//
// func (this *LRUCache) Put(key int, value int) {
// 	if node, ok := this.Cache[key]; ok {
// 		node.Val = value
// 		this.PutHead(node)
// 		return
// 	}
// 	if len(this.Cache) >= this.Capacity {
// 		this.Pop()
// 	}
// 	node := &LRUNode{Val: value, Key: key}
// 	this.Cache[key] = node
// 	this.PutHead(node)
// }
//
// func (this *LRUCache) Pop() {
// 	delete(this.Cache, this.LRU.Key)
// 	if this.LRU == this.MRU {
// 		this.LRU, this.MRU = nil, nil
// 		return
// 	}
// 	this.LRU = this.LRU.Prev
// }
//
// func (this *LRUCache) PutHead(node *LRUNode) {
// 	if this.MRU == nil {
// 		this.MRU = node
// 		this.LRU = node
// 		return
// 	}
// 	if node == this.MRU {
// 		return
// 	}
// 	if node == this.LRU {
// 		this.LRU = this.LRU.Prev
// 	}
// 	if node.Prev != nil {
// 		node.Prev.Next = node.Next
// 	}
// 	if node.Next != nil {
// 		node.Next.Prev = node.Prev
// 	}
// 	this.MRU.Prev = node
// 	node.Next = this.MRU
// 	this.MRU = node
// }
//
// /**
//  * Your LRUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);
//  */
