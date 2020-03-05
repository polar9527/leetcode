/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (36.77%)
 * Likes:    701
 * Dislikes: 856
 * Total Accepted:    243.3K
 * Total Submissions: 661.3K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 *
 *
 *
 *
 *
 */
package main

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)

	m := make(map[int]int)
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			if i-m[nums[i]] <= k {
				return true
			} else {
				m[nums[i]] = i
			}
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

// @lc code=end

func main() {
	a := []int{2, 2}
	containsNearbyDuplicate(a, 3)
}
