type MinStack struct {
    MinSta []int
    Sta []int
}


func Constructor() MinStack {
    this := MinStack{}
    return this
}


func (this *MinStack) Push(val int)  {
    if len(this.MinSta) == 0 || val <= this.MinSta[len(this.MinSta) - 1] {
        this.MinSta = append(this.MinSta, val)
    }
    this.Sta = append(this.Sta, val)
}


func (this *MinStack) Pop()  {
    val := this.Sta[len(this.Sta) - 1]
    if val == this.MinSta[len(this.MinSta) - 1] {
        this.MinSta = this.MinSta[:len(this.MinSta) - 1]
    }
    this.Sta = this.Sta[:len(this.Sta) - 1]
}


func (this *MinStack) Top() int {
    return this.Sta[len(this.Sta) - 1]
}


func (this *MinStack) GetMin() int {
    return this.MinSta[len(this.MinSta) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
