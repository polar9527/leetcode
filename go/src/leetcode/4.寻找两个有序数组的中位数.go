/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (37.36%)
 * Likes:    2659
 * Dislikes: 0
 * Total Accepted:    197.8K
 * Total Submissions: 520.3K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
 *
 * 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 *
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
package main

import (
	"math"
)

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	l, r := 0, m
	mediam1 := 0
	mediam2 := 0

	for l <= r {
		i := (l + r) / 2
		j := (m+n+1)/2 - i

		nums1_im1 := math.MinInt32
		if i != 0 {
			nums1_im1 = nums1[i-1]
		}
		nums1_i := math.MaxInt32
		if i != m {
			nums1_i = nums1[i]
		}

		nums2_jm1 := math.MinInt32
		if j != 0 {
			nums2_jm1 = nums2[j-1]
		}
		nums2_j := math.MaxInt32
		if j != n {
			nums2_j = nums2[j]
		}

		if nums1_im1 <= nums2_j {
			mediam1 = max(nums1_im1, nums2_jm1)
			mediam2 = min(nums1_i, nums2_j)
			l = i + 1
		} else {
			r = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(mediam1+mediam2) / 2.0
	}
	return float64(mediam1)
}

// @lc code=end
