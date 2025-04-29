package go_case

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (32.55%)
 * Likes:    6170
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
 * k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
 *
 * 你返回所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 解释：
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
 * 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
 * 注意，输出的顺序和三元组的顺序并不重要。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,1]
 * 输出：[]
 * 解释：唯一可能的三元组和不为 0 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,0,0]
 * 输出：[[0,0,0]]
 * 解释：唯一可能的三元组和为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
// func threeSum(nums []int) [][]int {
// 	var res [][]int
// 	sort.Ints(nums)

// 	for i := 0; i < len(nums)-2; i++ {
// 		n1 := nums[i]

// 		if nums[i] > 0 {
// 			return res
// 		}
// 		// n1 去重
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		l, r := i+1, len(nums)-1
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
// 			if n1+n2+n3 == 0 {
// 				//
// 				res = append(res, []int{n1, n2, n3})
// 				l++
// 				r--
// 				// b, c 去重
// 				for l < r && nums[l] == n2 {
// 					l++
// 				}
// 				for l < r && nums[r] == n3 {
// 					r--
// 				}
// 			} else if n1+n2+n3 > 0 {
// 				r--
// 			} else {
// 				l++
// 			}
// 		}
// 	}
// 	return res
// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i <= n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		j, k := i+1, n-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				jcur := nums[j]
				kcur := nums[k]
				res = append(res, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j] == jcur {
					j++
				}
				for j < k && nums[k] == kcur {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}

		}
	}
	return res
}

// @lc code=end
