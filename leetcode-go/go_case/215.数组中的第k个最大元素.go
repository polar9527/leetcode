package go_case

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
// func findKthLargest(nums []int, k int) int {
// 	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
// }

// func quickSelect(a []int, l, r, index int) int {
// 	q := randomPartition(a, l, r)
// 	if q == index {
// 		return a[q]
// 	} else if q < index {
// 		return quickSelect(a, q+1, r, index)
// 	}
// 	return quickSelect(a, l, q-1, index)
// }

// func randomPartition(a []int, l, r int) int {
// 	seed := time.Now().UnixNano()
// 	randomGenerator := rand.New(rand.NewSource(seed))
// 	i := randomGenerator.Intn(r-l+1) + l
// 	a[i], a[r] = a[r], a[i]
// 	return partition(a, l, r)
// }

// func partition(a []int, l, r int) int {
// 	// pivot 在 index r 处
// 	x := a[r]
// 	i := l - 1
// 	for j := l; j < r; j++ {
// 		if a[j] <= x {
// 			i++
// 			a[i], a[j] = a[j], a[i]
// 		}
// 	}
// 	a[i+1], a[r] = a[r], a[i+1]
// 	return i + 1
// }

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 小顶堆的顶部最小元素放在数组尾部，
		// 减小堆大小（数组长度），就相当于将最小元素删除
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		// 然后需要重新下沉新的顶部元素，将推恢复为小顶堆
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	// a[heapSize/2] 之后都是叶子，所以只需要对非叶子节点做下沉操作
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
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

// 利用 container/heap 库
// func findKthLargest(nums []int, k int) int {
// 	a := &kthLargest{k: k}
// 	for _, num := range nums {
// 		heap.Push(a, num)
// 		if a.Len() > a.k {
// 			heap.Pop(a)
// 		}
// 	}

// 	return a.IntSlice[0]
// }

// type kthLargest struct {
// 	sort.IntSlice
// 	k int
// }

// func (k *kthLargest) Pop() interface{} {
// 	a := k.IntSlice
// 	v := a[a.Len()-1]
// 	k.IntSlice = a[:k.Len()-1]
// 	return v
// }

// func (k *kthLargest) Push(v interface{}) {
// 	k.IntSlice = append(k.IntSlice, v.(int))
// }

// @lc code=end
