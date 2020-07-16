package main

type MinStack struct {
	A []int // save all data
	b []int // 保存所有非严格降序的元素
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.A = append(this.A, x)
	if len(this.b) == 0 || this.b[len(this.b)-1] >= x {
		this.b = append(this.b, x)
	}
}


func (this *MinStack) Pop()  {
	v := this.A[len(this.A)-1]
	this.A = this.A[0 : len(this.A)-1]
	if v == this.b[len(this.b)-1] {
		this.b = this.b[0 : len(this.b)-1]
	}
}


func (this *MinStack) Top() int {
    return this.A[len(this.A)-1]
}


func (this *MinStack) Min() int {
    return this.b[len(this.b)-1]
}