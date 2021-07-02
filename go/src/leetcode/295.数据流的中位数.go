/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (45.23%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    15K
 * Total Submissions: 33.1K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/
package leetcode

// @lc code=start
type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := make([]int, 0)
	max := make([]int, 0)

	return MedianFinder{min, max}
}

func (this *MedianFinder) AddNum(num int) {
	// 第一个元素，扔到小顶堆
	if len(this.minHeap) == 0 {
		this.minHeap = append(this.minHeap, num)
		return
	}

	if len(this.minHeap) < len(this.maxHeap) {
		// 添加到小顶堆
		if num < this.maxHeap[0] {
			this.minHeap = append(this.minHeap, this.maxHeap[0])
			this.minHeapUp()
			this.maxHeap[0] = num
			this.maxHeapDown()
		} else {
			this.minHeap = append(this.minHeap, num)
			this.minHeapUp()
		}
	} else {
		// 添加到大顶堆
		if num > this.minHeap[0] {
			this.maxHeap = append(this.maxHeap, this.minHeap[0])
			this.maxHeapUp()
			this.minHeap[0] = num
			this.minHeapDown()
		} else {
			this.maxHeap = append(this.maxHeap, num)
			this.maxHeapUp()
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.minHeap) == 0 && len(this.maxHeap) == 0 {
		return -1
	}
	if len(this.minHeap) == len(this.maxHeap) {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2.0
	} else if len(this.minHeap) < len(this.maxHeap) {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}

func (this *MedianFinder) minHeapUp() {
	child := len(this.minHeap) - 1
	parent := (len(this.minHeap) - 1) / 2
	for child > 0 && this.minHeap[parent] > this.minHeap[child] {
		this.minHeap[child], this.minHeap[parent] = this.minHeap[parent], this.minHeap[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) minHeapDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.minHeap) {
		min := childL
		if childR < len(this.minHeap) && this.minHeap[childL] > this.minHeap[childR] {
			min = childR
		}
		if this.minHeap[min] > this.minHeap[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.minHeap[min], this.minHeap[parent] = this.minHeap[parent], this.minHeap[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}

func (this *MedianFinder) maxHeapUp() {
	child := len(this.maxHeap) - 1
	parent := (len(this.maxHeap) - 1) / 2
	for child > 0 && this.maxHeap[parent] < this.maxHeap[child] {
		this.maxHeap[child], this.maxHeap[parent] = this.maxHeap[parent], this.maxHeap[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) maxHeapDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.maxHeap) {
		min := childL
		if childR < len(this.maxHeap) && this.maxHeap[childL] < this.maxHeap[childR] {
			min = childR
		}
		if this.maxHeap[min] < this.maxHeap[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.maxHeap[min], this.maxHeap[parent] = this.maxHeap[parent], this.maxHeap[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
