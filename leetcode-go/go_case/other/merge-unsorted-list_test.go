package go_case

import (
	"fmt"
	"testing"
)

// List 1:  3 -> 1 -> 4 -> 5 -> nil
// List 2:  2 -> 6 -> 1 -> 7 -> nil
// 输出
// List:  1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> nil

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// 打印链表（辅助函数，用于测试）
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}
func TestMergeListHeap(t *testing.T) {
	// 测试输入：List1: 3->1->4->5, List2: 2->6->1->7
	list1 := createList([]int{3, 1, 4, 5})
	list2 := createList([]int{2, 6, 1, 7})

	// 步骤2: 合并两个已排序链表
	// mergedList := mergeListHeap(list1, list2)
	mergedList := mergeListDivideAndConquer(list1, list2)

	// 打印输出
	fmt.Print("排序并合并后的链表: ")
	printList(mergedList) // 输出: 1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> nil
}
