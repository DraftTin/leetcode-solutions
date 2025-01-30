// package main
//
// type Node struct {
// 	Key  int
// 	Val  int
// 	Next *Node
// 	Prev *Node
// }
//
// type LRUCache struct {
// 	Cache    map[int]*Node
// 	Capacity int
// 	HeadNode *Node
// 	TailNode *Node
// }
//
// func Constructor(capacity int) LRUCache {
// 	this := LRUCache{Capacity: capacity}
// 	this.Cache = map[int]*Node{}
// 	return this
// }
//
// func (this *LRUCache) Get(key int) int {
// 	if this.Cache[key] == nil {
// 		return -1
// 	}
// 	node := this.Cache[key]
// 	this.PutTail(node)
// 	return node.Val
// }
//
// func (this *LRUCache) Put(key int, value int) {
// 	if this.Cache[key] == nil {
// 		if len(this.Cache) == this.Capacity {
// 			this.Pop()
// 		}
// 		node := &Node{Val: value, Key: key}
// 		this.Cache[key] = node
// 		this.PutTail(node)
// 	} else {
// 		node := this.Cache[key]
// 		node.Val = value
// 		this.PutTail(node)
// 	}
// }
//
// func (this *LRUCache) Pop() {
// 	if this.HeadNode == nil {
// 		return
// 	}
// 	delete(this.Cache, this.HeadNode.Key)
// 	this.HeadNode = this.HeadNode.Next
// }
//
// func (this *LRUCache) PutTail(node *Node) {
// 	if node == this.TailNode {
// 		return
// 	}
// 	if node.Prev != nil {
// 		node.Prev.Next = node.Next
// 	}
// 	if node.Next != nil {
// 		node.Next.Prev = node.Prev
// 	}
// 	if node == this.HeadNode {
// 		this.HeadNode = this.HeadNode.Next
// 	}
// 	if this.HeadNode == nil {
// 		this.HeadNode = node
// 		this.TailNode = node
// 	} else {
// 		this.TailNode.Next = node
// 		node.Prev = this.TailNode
// 		this.TailNode = node
// 	}
// }
//
// /**
//  * Your LRUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);
//  */
