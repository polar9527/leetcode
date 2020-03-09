/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 *
 * https://leetcode.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (20.48%)
 * Likes:    908
 * Dislikes: 944
 * Total Accepted:    112.5K
 * Total Submissions: 548.8K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * Given an array of integers, find out whether there are two distinct indices
 * i and j in the array such that the absolute difference between nums[i] and
 * nums[j] is at most t and the absolute difference between i and j is at most
 * k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3, t = 0
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1, t = 2
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,5,9,1,5,9], k = 2, t = 3
 * Output: false
 *
 *
 *
 *
 */
package main

import (
	"math"
)

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 1 || t < 0 {
		return false
	}
	numsInt64 := make([]int64, len(nums))
	for i, num := range nums {
		numsInt64[i] = int64(num) - int64(math.MinInt32)
	}

	// maintain K length window, in this window, all nums[i] (0<= i < k) will be put into bucket relatively
	buckets := make(map[int64]int64)

	for i, num := range numsInt64 {
		nth := num / (int64(t) + 1)
		if _, ok := buckets[nth]; ok {
			return true
		}
		if _, ok := buckets[nth-1]; ok && (num-buckets[nth-1] < int64(t)+1) {
			return true
		}
		if _, ok := buckets[nth+1]; ok && (buckets[nth+1]-num < int64(t)+1) {
			return true
		}
		buckets[nth] = num
		if i >= k {
			delete(buckets, numsInt64[i-k]/(int64(t)+1))
		}

	}

	return false
}

// @lc code=end
// func main() {
// 	a := []int{-3,3}
// 	ret := containsNearbyAlmostDuplicate(a, 2, 4)
// 	fmt.Println(ret)
// }
