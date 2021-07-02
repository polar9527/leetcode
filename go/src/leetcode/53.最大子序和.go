/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (51.00%)
 * Likes:    2024
 * Dislikes: 0
 * Total Accepted:    247K
 * Total Submissions: 481.7K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */
package leetcode

// @lc code=start
type Status struct {
	lSum int
	rSum int
	mSum int
	iSum int
}

func maxSubArray(nums []int) int {
	return getStatus(nums, 0, len(nums)-1).mSum
}

func getStatus(nums []int, l, r int) Status {
	if l == r {
		return Status{
			lSum: nums[l],
			rSum: nums[l],
			mSum: nums[l],
			iSum: nums[l],
		}
	}

	mi := (l + r) >> 1
	lStatus := getStatus(nums, l, mi)
	rStatus := getStatus(nums, mi+1, r)

	return pushUp(lStatus, rStatus)
}

func pushUp(l, r Status) Status {

	return Status{
		lSum: max(l.lSum, l.iSum+r.lSum),
		rSum: max(r.rSum, r.iSum+l.rSum),
		mSum: max(max(l.mSum, r.mSum), l.rSum+r.lSum),
		iSum: l.iSum + r.iSum,
	}
}

// func maxSubArray(nums []int) int {
// 	l := len(nums)
// 	if l == 0 {
// 		return -1
// 	}
// 	pre := nums[0]
// 	result := nums[0]
// 	for i := 1; i < l; i++ {
// 		pre = max(pre+nums[i], nums[i])
// 		result = max(result, pre)
// 	}
// 	return result
// }
//
// func max(x, y int) int {
//     if x < y {
//         return y
//     }
//     return x
// }
// @lc code=end
