package go_case

import "math"

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (43.19%)
 * Likes:    7521
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 3M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	if len(nums1) > len(nums2) {
// 		// m := len(nums1)
// 		// n := len(nums2)
// 		// j = [0,n-1]
// 		// j = (m+n+1)/2 > (n+n+1)/2 = n +1/2 会越界
// 		nums1, nums2 = nums2, nums1
// 	}
// 	m := len(nums1)
// 	n := len(nums2)
// 	nums1 = append([]int{math.MinInt32}, append(nums1, math.MaxInt32)...)
// 	nums2 = append([]int{math.MinInt32}, append(nums2, math.MaxInt32)...)
// 	i, j := 0, (m+n+1)/2
// 	for {
// 		if nums1[i] <= nums2[j+1] && nums1[i+1] >= nums2[j] {
// 			max1 := max(nums1[i], nums2[j])
// 			min2 := min(nums1[i+1], nums2[j+1])

// 			if (m+n)%2 > 0 {
// 				// 奇数， 前面已做处理， 即(m+n+1)/2
// 				return float64(max1)
// 			}
// 			return float64(max1+min2) / 2
// 		}
// 		i++
// 		j--
// 	}
// }

func findMedianSortedArrays(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a // 保证下面的 i 可以从 0 开始枚举
	}

	m, n := len(a), len(b)
	a = append([]int{math.MinInt}, append(a, math.MaxInt)...)
	b = append([]int{math.MinInt}, append(b, math.MaxInt)...)

	// 循环不变量：a[left-1] <= b[j+1]
	// 循环不变量：a[right+1] > b[j+1]
	// a 的长度是 m+2
	// 考虑到 i+1, 所以 right 最大 不能为 m+1，而是 m
	left, right := 0, m
	for left <= right { // 闭区间 [left, right] 不为空
		i := left + (right-left)>>1
		j := (m+n+1)/2 - i
		if a[i] <= b[j+1] {
			left = i + 1 // 缩小二分区间为 [i, right]
			// 更新后，left左边都是 a[i] <= b[j+1]
		} else {
			right = i - 1 // 缩小二分区间为 [left, i]
		}
	}

	// 此时 left 等于 right+1
	// a[left-1] <= b[j+1] 且 a[right+1] > b[j+1] = b[j]，所以答案是 i=left-1
	i := left - 1
	j := (m+n+1)/2 - i
	max1 := max(a[i], b[j])
	min2 := min(a[i+1], b[j+1])
	if (m+n)%2 > 0 {
		return float64(max1)
	}
	return float64(max1+min2) / 2
}

// @lc code=end
