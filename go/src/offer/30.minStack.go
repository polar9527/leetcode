package main

type MinStack struct {
	data []int
	m    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	d := make([]int, 0)
	m := make([]int, 0)
	return MinStack{d, m}
}

func (this *MinStack) Push(x int) {

	if len(this.m) == 0 {
		this.m = append(this.m, x)
	} else {
		if x < this.m[len(this.m)-1] {
			this.m = append(this.m, x)
		} else {
			this.m = append(this.m, this.m[len(this.m)-1])
		}
	}

	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	l := len(this.data)
	lm := len(this.m)
	if l != 0 {
		this.data = this.data[:l-1]
	}
	if lm != 0 {
		this.m = this.m[:lm-1]
	}
}

func (this *MinStack) Top() int {
	l := len(this.data)
	if l != 0 {
		return this.data[l-1]
	}
	panic("out of range")
}

func (this *MinStack) Min() int {
	lm := len(this.m)
	if lm != 0 {
		return this.m[lm-1]
	}
	panic("out of range")
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
