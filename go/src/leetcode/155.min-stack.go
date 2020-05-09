/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (40.97%)
 * Likes:    2790
 * Dislikes: 283
 * Total Accepted:    433.9K
 * Total Submissions: 1M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 *
 * Example:
 *
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 *
 *
 *
 *
 */
package main

// @lc code=start
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

func (this *MinStack) GetMin() int {
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
 * param_4 := obj.GetMin();
 */
// @lc code=end
