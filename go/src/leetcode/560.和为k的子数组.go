/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.82%)
 * Likes:    408
 * Dislikes: 0
 * Total Accepted:    46.1K
 * Total Submissions: 104.6K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的eenv

个数。
 *
 * 示例 1 :
 *
 *
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 *
 *
 * 说明 :
 *
 *
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 *
 *
*/
package main

import (
	"fmt"
)

// @lc code=start

func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	presum := 0
	count := 0

	// [0:i] 子串的和为k
	sumMap[0] = 1
	for i := 0; i < len(nums); i++ {
		// [j:i]
		presum += nums[i]
		if _, ok := sumMap[presum-k]; ok {
			count += sumMap[presum-k]
		}
		sumMap[presum] += 1
	}
	return count
}

// func subarraySum(nums []int, k int) int {
// 	l := len(nums)
// 	count := 0
// 	for i := 0; i < l; i++ {
// 		if nums[i] == k {
// 			count++
// 		}
// 		sum := nums[i]
// 		for j := i+1; j < l; j++ {
// 			sum += nums[j]
// 			if sum == k {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
// @lc code=end

func main() {
	nums := []int{1, 1, 1}
	k := 2
	ret := subarraySum(nums, k)
	fmt.Println(ret)
}
