/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode.cn/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (44.77%)
 * Likes:    2785
 * Dislikes: 0
 * Total Accepted:    936.9K
 * Total Submissions: 2.1M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
 *
 * 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j]
 * 处:
 *
 *
 * 0 <= j <= nums[i]
 * i + j < n
 *
 *
 * 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [2,3,0,1,4]
 * 输出: 2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 1000
 * 题目保证可以到达 nums[n-1]
 *
 *
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	curScope := 0
	maxScope := 0
	res := 0
	for i := 0; i <= curScope; i++ {
		if i+nums[i] > maxScope {
			maxScope = i + nums[i]
		}
		if i == curScope {
			curScope = maxScope
			res++
			if curScope >= len(nums)-1 {
				break
			}
		}
	}
	return res
}

func jump(nums []int) int {
	var res int
	nextRight, curRight := 0, 0

	for i, step := range nums {
		nextRight = max(nextRight, i+step)
		if i == curRight {
			curRight = nextRight
			res++
			if i == len(nums)-1 {
				res--
			}
		}
	}
	return res
}

// @lc code=end

