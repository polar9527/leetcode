package go_case

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
//  priority queue
// func mergeKLists(lists []*ListNode) *ListNode {

// 	h := hp{}
// 	for _, head := range lists {
// 		if head != nil {
// 			h = append(h, head)
// 		}
// 	}

// 	heap.Init(&h)
// 	dummy := &ListNode{}
// 	cur := dummy
// 	for h.Len() > 0 {
// 		curMinNode := heap.Pop(&h).(*ListNode)
// 		if curMinNode.Next != nil {
// 			heap.Push(&h, curMinNode.Next)
// 		}
// 		cur.Next = curMinNode
// 		cur = cur.Next
// 	}
// 	return dummy.Next
// }

// type hp []*ListNode

// func (h hp) Len() int           { return len(h) }
// func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆
// func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
// func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// merge and conquer
// 21. 合并两个有序链表（双指针）
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := ListNode{} // 用哨兵节点简化代码逻辑
	cur := &dummy       // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

// recursive
// func mergeKLists(lists []*ListNode) *ListNode {
// 	n := len(lists)
// 	if n == 0 {
// 		return nil
// 	}
// 	if n == 1 {
// 		return lists[0]
// 	}

// 	left := mergeKLists(lists[:n/2])
// 	right := mergeKLists(lists[n/2:])

// 	return mergeTwoLists(left, right)
// }

// iterative
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	for step := 1; step < n; step *= 2 {
		// n-step 的原因是 i 每次向前移动 2*step
		for i := 0; i < n-step; i += step * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+step])
		}
	}
	return lists[0]
}

// @lc code=end
