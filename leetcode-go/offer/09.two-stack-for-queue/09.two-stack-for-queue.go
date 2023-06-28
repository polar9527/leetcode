package offer

type CQueue struct {
	in  []int
	out []int
}

var fake = true

func Constructor() CQueue {
	return CQueue{make([]int, 0, 16), make([]int, 0, 16)}
}

func (this *CQueue) AppendTail(value int) {
	if fake {
		this.AppendTailFakeStack(value)
	} else {
		this.AppendTailStack(value)
	}
}

func (this *CQueue) DeleteHead() int {
	if fake {
		return this.DeleteHeadFakeStack()
	} else {
		return this.DeleteHeadStack()
	}
}

func (this *CQueue) AppendTailFakeStack(value int) {
	iLen := len(this.in)
	oLen := len(this.out)
	if oLen == 0 && iLen == cap(this.in) {
		this.out = this.in
		this.in = make([]int, 0, 16)
	}
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHeadFakeStack() int {
	iLen := len(this.in)
	oLen := len(this.out)
	head := -1
	if oLen == 0 && iLen == 0 {
		return head
	}
	if oLen != 0 {
		head = this.out[0]
		if oLen > 1 {
			this.out = this.out[1:]
		} else if iLen == cap(this.in) {
			this.out = this.in
			this.in = make([]int, 0, 16)
		} else {
			this.out = make([]int, 0, 16)
		}
		return head
	}
	head = this.in[0]
	if iLen == cap(this.in) {
		this.out = this.in[1:]
		this.in = make([]int, 0, 16)
	} else {
		this.in = this.in[1:]
	}
	return head
}

func (this *CQueue) AppendTailStack(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHeadStack() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
		}
		this.in = []int{}
	}
	lastIndex := len(this.out) - 1
	lastItem := this.out[lastIndex]
	if lastIndex == 0 {
		this.out = []int{}
		return lastItem
	}
	this.out = this.out[:lastIndex]
	return lastItem
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
