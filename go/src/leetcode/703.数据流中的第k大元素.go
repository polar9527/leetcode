/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (43.84%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    39.7K
 * Total Submissions: 81.9K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * 设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
 *
 * 请实现 KthLargest 类：
 *
 *
 * KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
 * int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 * 输出：
 * [null, 4, 5, 5, 8, 8]
 *
 * 解释：
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3);   // return 4
 * kthLargest.add(5);   // return 5
 * kthLargest.add(10);  // return 5
 * kthLargest.add(9);   // return 8
 * kthLargest.add(4);   // return 8
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * -10^4
 * -10^4
 * 最多调用 add 方法 10^4 次
 * 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
 *
 *
*/

// @lc code=start
type minHeap struct {
	k    int
	heap []int
}

func createMinHeap(k int, nums []int) *minHeap {
	heap := &minHeap{
		k:    k,
		heap: []int{},
	}
	for _, v := range nums {
		heap.add(v)
	}
	return heap
}

func (this *minHeap) add(val int) {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		this.up(len(this.heap) - 1)
	} else if val > this.heap[0] {
		this.heap[0] = val
		this.down(0)
	}
}

func (this *minHeap) down(i int) {
	for 2*i+1 < len(this.heap) {
		child := 2*i + 1
		if child+1 < len(this.heap) && this.heap[child+1] < this.heap[child] {
			child++
		}
		if this.heap[child] < this.heap[i] {
			this.heap[i], this.heap[child] = this.heap[child], this.heap[i]
			i = child
		} else {
			break
		}
	}
}

func (this *minHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) >> 1
		if this.heap[parent] > this.heap[i] {
			this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent]
			i = parent
		} else {
			break
		}
	}
}

type KthLargest struct {
	heap *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{heap: createMinHeap(k, nums)}
}

func (this *KthLargest) Add(val int) int {
	this.heap.add(val)
	return this.heap.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

