/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (52.44%)
 * Likes:    2980
 * Dislikes: 210
 * Total Accepted:    529.7K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 *
 *
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 */
package main

// @lc code=start
func findKthLargest(nums []int, k int) int {

	quickSelect(nums, k-1)
	return nums[k-1]
	// ret := heapSelect(nums, k)
	// return ret
}

// 32/32 cases passed (12 ms)
// Your runtime beats 44.76 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3.5 MB)
func quickSelect(nums []int, k int) {
	l := len(nums)
	for lo, hi := 0, l-1; lo < hi; {
		i, j := lo, hi
		// rand.Seed(time.Now().UnixNano())
		// ir := rand.Intn(hi-lo)+lo
		// nums[lo], nums[ir] = nums[lo],nums[ir]
		pivot := nums[lo]
		for i < j {
			for i < j && pivot >= nums[j] {
				j--
			}
			nums[i] = nums[j]
			for i < j && pivot <= nums[i] {
				i++
			}
			nums[j] = nums[i]
		}
		nums[i] = pivot
		if k <= i {
			hi = i - 1
		}
		if i <= k {
			lo = i + 1
		}
	}

}

// 32/32 cases passed (4 ms)
// Your runtime beats 99.37 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3.5 MB)
func heapSelect(nums []int, k int) int {
	l := len(nums)

	// create (k) min heap
	startK := k / 2
	for start := startK; start >= 0; start-- {
		minHeap(start, k-1, nums[:k])
	}

	if k != l {
		// create (n - k) max heap
		startNK := (l - k) / 2
		for start := startNK; start >= 0; start-- {
			maxHeap(start, l-k-1, nums[k:])
		}

		for nums[0] < nums[k] {
			nums[0], nums[k] = nums[k], nums[0]

			startK := 0
			minHeap(startK, k-1, nums[:k])

			startNK := 0
			maxHeap(startNK, l-k-1, nums[k:])
		}
	}

	ret := nums[0]
	return ret
}

func maxHeap(start, end int, array []int) {
	root := start
	for true {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 <= end && array[child] < array[child+1] {
			child++
		}
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			root = child
		} else {
			break
		}
	}
}

func minHeap(start, end int, array []int) {
	root := start
	for true {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 <= end && array[child] > array[child+1] {
			child++
		}
		if array[root] > array[child] {
			array[root], array[child] = array[child], array[root]
			root = child
		} else {
			break
		}
	}
}

// @lc code=end

// func main(){
// 	// a := []int{-1,2,0}
// 	// a := []int{3,2,1,5,6,4}
// 	a := []int{3,2,3,1,2,4,5,5,6}
// 	// a := []int{1}
//
//
// 	ret := findKthLargest(a, 4)
// 	fmt.Println(a)
// 	fmt.Println(ret)
// }
