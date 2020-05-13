package main

// import (
// 	"fmt"
// )

type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor() MaxQueue {
	q := make([]int, 0)
	m := make([]int, 0)

	return MaxQueue{q, m}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) != 0 && len(this.max) != 0 {
		return this.queue[this.max[0]]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.max) != 0 && this.queue[this.max[len(this.max)-1]] <= value {
		this.max = this.max[:len(this.max)-1]
	}
	this.queue = append(this.queue, value)
	this.max = append(this.max, len(this.queue)-1)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 || len(this.max) == 0 {
		return -1
	}
	// 修改m
	if this.max[0] == 0 {
		this.max = this.max[1:]
	}
	// 修改q
	ret := this.queue[0]
	this.queue = this.queue[1:]
	for i := range this.max {
		this.max[i]--
	}
	return ret
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

// func main(){
// 	mq := new(MaxQueue)
// 	fmt.Println(mq.Pop_front())
// 	fmt.Println(mq.Pop_front())
// 	fmt.Println(mq.Pop_front())
// 	fmt.Println(mq.Pop_front())
// 	fmt.Println(mq.Pop_front())
// 	mq.Push_back(15)
// 	fmt.Println(mq.Max_value())
// 	mq.Push_back(9)
// 	fmt.Println(mq.Max_value())
//
// }
