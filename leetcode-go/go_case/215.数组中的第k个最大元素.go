package go_case

import (
	"container/heap"
	"math/rand"
	"sort"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (63.63%)
 * Likes:    2252
 * Dislikes: 0
 * Total Accepted:    906.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4], k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6], k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Intn(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	// pivot 在 index r 处
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

// 堆排序
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeapSize(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 大顶堆的顶部最大元素放在数组尾部，
		// 然后减小堆大小（数组长度），就相当于将刚才置换到堆尾的最大元素删除
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		// 然后需要重新下沉新的顶部的小元素，将堆恢复为大顶堆
		maxHeapify(nums, 0, heapSize)
	} // 经过 k-1 次迭代之后，堆顶部就是第k大的元素
	return nums[0]
}

func buildMaxHeapSize(a []int, heapSize int) {
	// a[heapSize/2] 之后都是叶子，所以只需要对非叶子节点做下沉操作
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	lc, rc := i*2+1, i*2+2
	largest := i
	if lc < heapSize && a[lc] > a[largest] {
		largest = lc
	}
	if rc < heapSize && a[rc] > a[largest] {
		largest = rc
	}
	if largest != i {
		// 如果 i 不是最大的，那么就要将a[i] 下沉到左孩子或者右孩子当中去
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// 利用 container/heap 库
func findKthLargest(nums []int, k int) int {
	a := &kthLargest{k: k}
	for _, num := range nums {
		heap.Push(a, num)
		if a.Len() > a.k {
			heap.Pop(a)
		}
	}

	return a.IntSlice[0]
}

type kthLargest struct {
	sort.IntSlice
	k int
}

func (k *kthLargest) Pop() interface{} {
	a := k.IntSlice
	v := a[a.Len()-1]
	k.IntSlice = a[:k.Len()-1]
	return v
}

func (k *kthLargest) Push(v interface{}) {
	k.IntSlice = append(k.IntSlice, v.(int))
}

// 快排
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 这个分区法比较低效，尤其是 在处理 有很多相等元素的时候
	var partition func(int, int) int
	partition = func(l, r int) int {
		// 0,1,2,3,4,5
		// l = 2, r = 5
		// r-l+1 = 4
		// [l,4+l)
		// [2,6) -> [2,5]
		pivot := rand.Intn(r-l+1) + l
		pivotVal := nums[pivot]
		nums[r], nums[pivot] = nums[pivot], nums[r]

		less := l - 1
		for more := l; more < r; more++ {
			if nums[more] <= pivotVal {
				less++
				nums[less], nums[more] = nums[more], nums[less]
			}
		}
		nums[less+1], nums[r] = nums[r], nums[less+1]
		return less + 1
	}

	var quickSelect func(int, int, int) int
	quickSelect = func(l, r, index int) int {

		pivot := partition(l, r)
		if pivot == index {
			return nums[index]
		}
		if pivot < index {
			return quickSelect(pivot+1, r, index)
		}
		// pivot > index
		return quickSelect(l, pivot-1, index)
	}

	return quickSelect(0, n-1, n-k)
}

// 快排， Hoare 分区
func findKthLargest(nums []int, k int) int {
	n := len(nums)

	var quickselect func(int, int, int) int
	quickselect = func(l, r, index int) int {
		//  l <= index <= r
		if l == r {
			return nums[index]
		}
		partition := nums[l]
		i := l - 1
		j := r + 1
		for i < j {
			for i++; nums[i] < partition; i++ {
			}
			for j--; nums[j] > partition; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if index <= j {
			return quickselect(l, j, index)
		} else {
			return quickselect(j+1, r, index)
		}
	}
	return quickselect(0, n-1, n-k)
}

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	var maxHeapify func(int, int)
	maxHeapify = func(index, hs int) {
		// if index >= hs {
		// 	return
		// }
		lc, rc := 2*index+1, 2*index+2
		largest := index
		if lc < hs && nums[lc] > nums[largest] {
			largest = lc
		}
		// 父节点必须大于两个子节点
		if rc < hs && nums[rc] > nums[largest] {
			largest = rc
		}

		if largest != index {
			nums[index], nums[largest] = nums[largest], nums[index]
			maxHeapify(largest, hs)
		}
	}

	buildMaxHeap := func(hs int) {
		for i := hs / 2; i >= 0; i-- {
			maxHeapify(i, hs)
		}
	}

	buildMaxHeap(heapSize)
	// 0,1,2,3,4,5,6
	// k = 2
	// 要从大顶堆 堆顶取出 k-1 个 元素， 放在数组末尾
	// assert heapSize >= k
	for i := 1; i <= k-1; i++ {
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		maxHeapify(0, heapSize)
	}
	return nums[0]
}

// @lc code=end
