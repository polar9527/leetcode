package go_case

import "container/heap"

type ListNode struct {
	Next *ListNode
	Val  int
}

// 排序链表（归并排序）
func sortList(head *ListNode) *ListNode {
	// 基本情况：链表为空或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}
	// 使用快慢指针找到中间节点
	mid := findMid(head)
	right := mid.Next
	mid.Next = nil // 断开链表为两部分
	// 递归排序左半部分和右半部分
	leftSorted := sortList(head)
	rightSorted := sortList(right)
	// 合并两个已排序链表
	return mergeTwoLists(leftSorted, rightSorted)
}

// 快慢指针找中间节点（慢指针为中间节点）
func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 合并两个已排序链表（迭代法）
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 虚拟头节点，简化操作
	current := dummy
	// 遍历两个链表，比较节点值
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	// 处理剩余节点
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummy.Next
}

func mergeListDivideAndConquer(list1, list2 *ListNode) *ListNode {
	sortedList1 := sortList(list1)
	sortedList2 := sortList(list2)

	// 步骤2: 合并两个已排序链表
	return mergeTwoLists(sortedList1, sortedList2)
}

func mergeListHeap(list1, list2 *ListNode) *ListNode {
	// var hp1, hp2 *hp
	hp1 := &hp{
		lt: make([]*ListNode, 0),
	}
	hp2 := &hp{
		lt: make([]*ListNode, 0),
	}

	head1 := list1
	for head1 != nil {
		heap.Push(hp1, head1)
		head1 = head1.Next
	}

	head2 := list2
	for head2 != nil {
		heap.Push(hp2, head2)
		head2 = head2.Next
	}

	dummy := &ListNode{}
	pre := dummy
	head := dummy.Next
	for hp1.Len() > 0 && hp2.Len() > 0 {
		var cur *ListNode
		if hp1.lt[0].Val < hp2.lt[0].Val {
			cur = heap.Pop(hp1).(*ListNode)
		} else {
			cur = heap.Pop(hp2).(*ListNode)
		}
		head = cur
		pre.Next = head

		pre = head
		head = head.Next
	}
	for hp1.Len() > 0 {
		head = heap.Pop(hp1).(*ListNode)
		pre.Next = head

		pre = head
		head = head.Next
	}

	for hp2.Len() > 0 {
		head = heap.Pop(hp2).(*ListNode)
		pre.Next = head

		pre = head
		head = head.Next
	}

	return dummy.Next
}

type hp struct {
	lt []*ListNode
}

func (h hp) Len() int {
	return len(h.lt)
}
func (h hp) Less(i, j int) bool {
	return h.lt[i].Val < h.lt[j].Val
}

func (h hp) Swap(i, j int) {
	h.lt[i], h.lt[j] = h.lt[j], h.lt[i]
}

func (h *hp) Push(item interface{}) {
	h.lt = append(h.lt, item.(*ListNode))
}

func (h *hp) Pop() any {
	old := h.lt
	res := old[len(old)-1]
	h.lt = old[:len(old)-1]
	return res
}
