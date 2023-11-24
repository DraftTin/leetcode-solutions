type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] < a[j] }

func (h *IntHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	size := len(old)
	*h = old[:size-1]
	return old[size-1]
}

func findKthLargest(nums []int, k int) int {
	h := make(IntHeap, k)
	copy(h, nums[:k])
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if nums[i] > h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}
	return h[0]
}
