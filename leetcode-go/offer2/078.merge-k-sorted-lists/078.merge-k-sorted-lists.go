package offer2

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (58.32%)
 * Likes:    2616
 * Dislikes: 0
 * Total Accepted:    699.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var h minHeap
	for _, listHeadNode := range lists {
		if listHeadNode != nil {
			heap.Push(&h, &nodePair{v: listHeadNode.Val, n: listHeadNode})
		}
	}
	dumpHead := &ListNode{}
	temp := dumpHead
	for len(h) > 0 {
		top := heap.Pop(&h).(*nodePair)
		temp.Next = top.n
		temp = temp.Next
		if top.n.Next != nil {
			heap.Push(&h, &nodePair{v: top.n.Next.Val, n: top.n.Next})
		}
	}
	return dumpHead.Next
}

type nodePair struct {
	v int
	n *ListNode
}

type minHeap []*nodePair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*nodePair))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
