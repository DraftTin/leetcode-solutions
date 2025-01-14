type PriorQueue []int

func (h PriorQueue) Len() int           { return len(h) }
func (h PriorQueue) Less(i, j int) bool { return h[i] < h[j] }
func (h PriorQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorQueue) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *PriorQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
  Queue *PriorQueue
}


func Constructor(n int) SeatManager {
  sm := SeatManager{}
  que := make(PriorQueue, n)
  sm.Queue = &que
  for i := 0; i < n; i++ {
    (*sm.Queue)[i] = i + 1
  }
  heap.Init(sm.Queue)
  return sm
}


func (this *SeatManager) Reserve() int {
  number := (*this.Queue)[0]
  heap.Pop(this.Queue)
  return number
}


func (this *SeatManager) Unreserve(seatNumber int)  {
  heap.Push(this.Queue, seatNumber)
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
