/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.61%)
 * Likes:    2031
 * Dislikes: 0
 * Total Accepted:    669.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		// 剪枝
		if n1 > target && target >= 0 {
			break
		}
		// 去重 n1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			// 剪枝
			if n1+n2 > target && target >= 0 {
				break
			}
			// 去重 n2
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				// 剪枝
				if n1+n2+n3 > target && target >= 0 {
					break
				}
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					l++
					r--
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if n1+n2+n3+n4 < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}

// @lc code=end

