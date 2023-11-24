type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IntHeap) Push(num interface{}) {
	(*h) = append((*h), num.(int))
}

func (h *IntHeap) Pop() interface{} {
	num := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return num
}

type KthLargest struct {
	H *IntHeap
	K int
}

func Constructor(k int, nums []int) KthLargest {
	var ins KthLargest
	if len(nums) < k {
		ins = KthLargest{(*IntHeap)(&nums), k}
	} else {
		kNums := nums[:k]
		ins = KthLargest{(*IntHeap)(&kNums), k}
	}
	heap.Init(ins.H)
	for i := k; i < len(nums); i++ {
		ins.Add(nums[i])
	}
	return ins
}

func (this *KthLargest) Add(val int) int {
	// maintain an arr with size-k
	if this.H.Len() < this.K {
		heap.Push(this.H, val)
	} else if (*this.H)[0] < val {
		heap.Pop(this.H)
		heap.Push(this.H, val)
	}
	return (*this.H)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
