package go_case

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode.cn/problems/sort-list/description/
 *
 * algorithms
 * Medium (65.66%)
 * Likes:    2093
 * Dislikes: 0
 * Total Accepted:    426.2K
 * Total Submissions: 649.1K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5 <= Node.val <= 10^5
 *
 *
 *
 *
 * 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
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
// recursive
// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	// 876. 链表的中间结点（快慢指针）, 修订版本，增加断开逻辑
// 	middleNode := func(head *ListNode) *ListNode {
// 		pre, slow, fast := head, head, head
// 		for fast != nil && fast.Next != nil {
// 			pre = slow // 记录 slow 的前一个节点
// 			slow = slow.Next
// 			fast = fast.Next.Next
// 		}

// 		pre.Next = nil // 断开逻辑：断开 slow 的前一个节点和 slow 的连接

// 		return slow
// 	}

// 	// 21. 合并两个有序链表（双指针）
// 	mergeTwoLists := func(list1, list2 *ListNode) *ListNode {
// 		dummy := ListNode{} // 用哨兵节点简化代码逻辑
// 		cur := &dummy       // cur 指向新链表的末尾
// 		for list1 != nil && list2 != nil {
// 			if list1.Val < list2.Val {
// 				cur.Next = list1 // 把 list1 加到新链表中
// 				list1 = list1.Next
// 			} else { // 注：相等的情况加哪个节点都是可以的
// 				cur.Next = list2 // 把 list2 加到新链表中
// 				list2 = list2.Next
// 			}
// 			cur = cur.Next
// 		}
// 		// 拼接剩余链表
// 		if list1 != nil {
// 			cur.Next = list1
// 		} else {
// 			cur.Next = list2
// 		}
// 		return dummy.Next
// 	}

// 	head2 := middleNode(head) // head2 的前一个节点和 head2 的连接 已经断开
// 	head = sortList(head)
// 	head2 = sortList(head2)
// 	return mergeTwoLists(head, head2)
// }

func sortList(head *ListNode) *ListNode {
	getListlength := func(head *ListNode) int {
		length := 0
		for head != nil {
			length++
			head = head.Next
		}
		return length
	}

	splitList := func(head *ListNode, size int) *ListNode {
		// 遍历到 cur 指向 指定长度链表中 的 倒数第一个节点
		cur := head
		for i := 0; i < size-1 && cur != nil; i++ {
			cur = cur.Next
		}

		if cur == nil || cur.Next == nil {
			// cur 此时指向 当前链表的第一个子链表的最后一个节点
			// cur == nil， (head, cur] 这个左开右闭区间上， 要么 长度 < size-1
			// 即便 循环走完了 size-1 步，(head, cur] 这个左开右闭区间上 cur == nil，也意味着区间长度 < size-1
			// 意味着 head 代表的链表长度 连划分第一个长度为 size 的 子链表都不够
			return nil
		}
		// 此时 cur 已经往前走了 size - 1 步，[head, cur] 这个闭区间内有 size 个节点

		nextHead := cur.Next
		cur.Next = nil // 切断两个子链表的连接
		return nextHead
	}

	// 21. 合并两个有序链表（双指针）, 修订版， 增加返回链表为节点信息
	mergeTwoLists := func(list1, list2 *ListNode) (*ListNode, *ListNode) {
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

		// 由于 list1 或者 list2 可能还有剩余，cur 当前可能并没有指向末尾节点，需要继续往后遍历
		for cur.Next != nil {
			cur = cur.Next
		}

		return dummy.Next, cur
	}

	length := getListlength(head) // 获取链表长度
	dummy := &ListNode{Next: head}

	// 分治归并算法, step 为参与merge的链表长度
	for step := 1; step < length; step *= 2 {
		nextListTail := dummy
		cur := dummy.Next
		// 每两个长度为 step 的子链表 merge
		for cur != nil {

			head1 := cur
			head2 := splitList(head1, step)
			// 下一对 子链表 的 起始节点
			cur = splitList(head2, step)

			// 合并这对链表
			head, tail := mergeTwoLists(head1, head2)

			// 合并后的链表头节点 head 接入 最终的排序链表的尾部
			nextListTail.Next = head
			nextListTail = tail
		}
	}
	return dummy.Next
}

// @lc code=end
