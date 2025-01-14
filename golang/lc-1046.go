type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] > a[j] }

func (h *IntHeap) Push(num interface{}) {
	(*h) = append(*h, num.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	size := len(old)
	*h = old[:size-1]
	return old[size-1]
}

func lastStoneWeight(stones []int) int {
	h := make(IntHeap, len(stones))
	copy(h, stones)
	heap.Init(&h)
	for len(h) != 0 && len(h) != 1 {
		max1 := h[0]
		heap.Pop(&h)
		max2 := h[0]
		heap.Pop(&h)
		if max1 > max2 {
			max1 = max1 - max2
			heap.Push(&h, max1)
			continue
		}
	}
	if len(h) == 0 {
		return 0
	}
	return h[0]
}
