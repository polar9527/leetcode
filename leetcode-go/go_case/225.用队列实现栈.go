/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 *
 * https://leetcode.cn/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (65.59%)
 * Likes:    940
 * Dislikes: 0
 * Total Accepted:    439.5K
 * Total Submissions: 668.7K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
 *
 * 实现 MyStack 类：
 *
 *
 * void push(int x) 将元素 x 压入栈顶。
 * int pop() 移除并返回栈顶元素。
 * int top() 返回栈顶元素。
 * boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty
 * 这些操作。
 * 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyStack", "push", "push", "top", "pop", "empty"]
 * [[], [1], [2], [], [], []]
 * 输出：
 * [null, null, null, 2, 2, false]
 *
 * 解释：
 * MyStack myStack = new MyStack();
 * myStack.push(1);
 * myStack.push(2);
 * myStack.top(); // 返回 2
 * myStack.pop(); // 返回 2
 * myStack.empty(); // 返回 False
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= x <= 9
 * 最多调用100 次 push、pop、top 和 empty
 * 每次调用 pop 和 top 都保证栈不为空
 *
 *
 *
 *
 * 进阶：你能否仅用一个队列来实现栈。
 *
 */

// @lc code=start

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{
		q1: make([]int, 0),
		q2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	this.Move()
}

// 通过 将 q1 中的入队数据 重新入队 q2 使得 最新push 进入 的数据 排在 q2 最前面，
// 从而在pop的时候最先pop出来
func (this *MyStack) Move() {
	if len(this.q1) == 0 {
		this.q1, this.q2 = this.q2, this.q1
	} else {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
		this.Move()
	}
}

func (this *MyStack) Pop() int {
	v := this.q1[0]
	this.q1 = this.q1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

