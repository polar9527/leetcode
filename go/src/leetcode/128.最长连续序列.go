/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (48.63%)
 * Likes:    391
 * Dislikes: 0
 * Total Accepted:    54K
 * Total Submissions: 107.4K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */
package leetcode

// @lc code=start
func longestConsecutive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	hashMap := make(map[int]bool)
	ans := 1
	for _, num := range nums {
		if _, ok := hashMap[num]; !ok {
			hashMap[num] = true
		}
	}

	for k := range hashMap {
		if _, ok := hashMap[k-1]; !ok {
			count := 1
			for i := k + 1; ; {
				if _, ok := hashMap[i]; ok {
					i++
					count++
				} else {
					if count > ans {
						ans = count
					}
					break
				}
			}
		}
	}
	return ans
}

// @lc code=end
