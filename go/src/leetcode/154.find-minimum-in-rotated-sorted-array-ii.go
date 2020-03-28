/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (40.26%)
 * Likes:    701
 * Dislikes: 192
 * Total Accepted:    162.9K
 * Total Submissions: 404.2K
 * Testcase Example:  '[1,3,5]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 *
 * Find the minimum element.
 *
 * The array may contain duplicates.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5]
 * Output: 1
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,0,1]
 * Output: 0
 *
 * Note:
 *
 *
 * This is a follow up problem to Find Minimum in Rotated Sorted Array.
 * Would allow duplicates affect the run-time complexity? How and why?
 *
 *
 */
package main

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	lo := 0
	hi := len(nums) - 1
	mi := lo
	for nums[lo] >= nums[hi] {
		if lo+1 == hi {
			mi = hi
			break
		}
		mi = (lo + hi) / 2
		if nums[mi] == nums[hi] && nums[mi] == nums[lo] {
			min := nums[lo]
			for i := lo; i <= hi; i++ {
				if nums[i] < min {
					min = nums[i]
				}
			}
			return min
		}
		if nums[mi] <= nums[hi] {
			hi = mi
		} else if nums[mi] >= nums[lo] {
			lo = mi
		}
	}
	return nums[mi]
}

// @lc code=end
