package main

type CQueue struct {
	Data []int
}

func Constructor() CQueue {
	d := make([]int, 0)
	return CQueue{d}
}

func (this *CQueue) AppendTail(value int) {
	this.Data = append(this.Data, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Data) == 0 {
		return -1
	}
	ret := this.Data[0]
	this.Data = append(this.Data[:0], this.Data[1:]...)
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
