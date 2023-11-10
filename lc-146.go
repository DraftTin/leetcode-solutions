type Node struct {
  Key int
  Val int
  Next *Node
  Prev *Node
}

type LRUCache struct {
  Cache map[int]*Node
  Capacity int
  LRUNode *Node
  TailNode *Node
}


func Constructor(capacity int) LRUCache {
  this := LRUCache{Capacity: capacity}
  this.Cache = map[int]*Node{}
  return this
}


func (this *LRUCache) Get(key int) int {
  fmt.Println("get")
  if this.Cache[key] == nil {
    return -1
  }
  node := this.Cache[key]
  this.PutTail(node)
  return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
  fmt.Println("put")
  if this.Cache[key] == nil {
    if len(this.Cache) == this.Capacity {
      this.Pop()
    }
    node := &Node{Val: value, Key: key}
    this.Cache[key] = node
    this.PutTail(node)
  } else {
    node := this.Cache[key]
    node.Val = value
    this.PutTail(node)
  }
}

func (this *LRUCache) Pop() {
  if this.LRUNode == this.TailNode {
    this.TailNode = nil
  }
  fmt.Println(this.LRUNode)
  delete(this.Cache, this.LRUNode.Key)
  this.LRUNode = this.LRUNode.Next
}

func (this *LRUCache) PutTail(node *Node) {
  if node == this.TailNode {
    return
  }
  if node.Prev != nil {
    node.Prev.Next = node.Next
  }
  if node.Next != nil {
    node.Next.Prev = node.Prev
  }
  if node == this.LRUNode {
    this.LRUNode = node.Next
  }
  if this.TailNode == nil {
    this.TailNode = node
    this.LRUNode = node
  } else {
    this.TailNode.Next = node
    node.Prev = this.TailNode
    this.TailNode = node
  }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
