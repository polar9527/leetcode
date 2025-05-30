package go_case

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode.cn/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (56.27%)
 * Likes:    1126
 * Dislikes: 0
 * Total Accepted:    217.5K
 * Total Submissions: 381.8K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
 *
 *
 * 例如 arr = [2,3,4] 的中位数是 3 。
 * 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
 *
 *
 * 实现 MedianFinder 类:
 *
 *
 *
 * MedianFinder() 初始化 MedianFinder 对象。
 *
 *
 * void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
 *
 *
 * double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10^-5 以内的答案将被接受。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
 * [[], [1], [2], [], [3], []]
 * 输出
 * [null, null, null, 1.5, null, 2.0]
 *
 * 解释
 * MedianFinder medianFinder = new MedianFinder();
 * medianFinder.addNum(1);    // arr = [1]
 * medianFinder.addNum(2);    // arr = [1, 2]
 * medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
 * medianFinder.addNum(3);    // arr[1, 2, 3]
 * medianFinder.findMedian(); // return 2.0
 *
 * 提示:
 *
 *
 * -10^5 <= num <= 10^5
 * 在调用 findMedian 之前，数据结构中至少有一个元素
 * 最多 5 * 10^4 次调用 addNum 和 findMedian
 *
 *
*/

// @lc code=start
type MedianFinder struct {
	left  hp // 入堆的元素取相反数，变成最大堆
	right hp // 最小堆
}

func Constructor() (_ MedianFinder) { return }

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == mf.right.Len() {
		heap.Push(&mf.right, num)
		heap.Push(&mf.left, -heap.Pop(&mf.right).(int))
	} else {
		heap.Push(&mf.left, -num)
		heap.Push(&mf.right, -heap.Pop(&mf.left).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64(-mf.left.IntSlice[0])
	}
	return float64(mf.right.IntSlice[0]-mf.left.IntSlice[0]) / 2.0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
