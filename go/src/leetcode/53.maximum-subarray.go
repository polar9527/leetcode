/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax

	if nums == nil || len(nums) <= 0 {
		return -1
	}

	ret := intMin
	cur_sum := 0
	for _, v := range nums {
		if cur_sum <= 0 {
			cur_sum = v
		} else {
			cur_sum += v
		}

		if cur_sum > ret {
			ret = cur_sum
		}
	}
	return ret
}
