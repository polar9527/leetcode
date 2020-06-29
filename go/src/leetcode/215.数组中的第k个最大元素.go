/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (62.59%)
 * Likes:    548
 * Dislikes: 0
 * Total Accepted:    143.2K
 * Total Submissions: 224.6K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */
package main

import (
	"math/rand"
	"time"
)

// @lc code=start
func findKthLargestHeapSort(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	// |<-          l         ->|
	//  ------------------------
	// |<- l-k+1 ->| <-  k-1  ->|
	// |0... ...l-k|l-k+1... l-1|
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	// 数组尾部元素 a[heapSize-1] 的父节点是a[(heapSize-1)/2],
	// 这个父节点是最后一个带叶子的节点，从这里开始遍历所有的带叶子节点
	// 使得每个带叶子的节点符合堆的性质
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

func findKthLargestQuickSelect(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(list []int, l, r, index int) int {
	pivot := randPartition(list, l, r)
	if pivot == index {
		return list[pivot]
	} else if pivot < index {
		return quickSelect(list, pivot+1, r, index)
	} else {
		return quickSelect(list, l, pivot-1, index)
	}
}

func randPartition(list []int, l, r int) int {
	i := rand.Intn(r-l+1) + l
	list[i], list[r] = list[r], list[i]
	return partition(list, l, r)
}

func partition(list []int, l, r int) int {
	pivot := list[r]
	// s smaller 指向小于等于 pivot 的元素的index
	s := l - 1
	// g greater 指向大于 pivot 的元素的index
	for g := l; g < r; g++ {
		if list[g] <= pivot {
			s++
			list[s], list[g] = list[g], list[s]
		}
	}
	list[s+1], list[r] = list[r], list[s+1]
	return s + 1
}

// @lc code=end
