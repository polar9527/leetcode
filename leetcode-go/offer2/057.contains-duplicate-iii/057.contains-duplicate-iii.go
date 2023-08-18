package offer2

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode.cn/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Hard (30.10%)
 * Likes:    704
 * Dislikes: 0
 * Total Accepted:    95.9K
 * Total Submissions: 318K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。
 *
 * 找出满足下述条件的下标对 (i, j)：
 *
 *
 * i != j,
 * abs(i - j) <= indexDiff
 * abs(nums[i] - nums[j]) <= valueDiff
 *
 *
 * 如果存在，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
 * 输出：true
 * 解释：可以找出 (i, j) = (0, 3) 。
 * 满足下述 3 个条件：
 * i != j --> 0 != 3
 * abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
 * abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
 * 输出：false
 * 解释：尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 1 <= indexDiff <= nums.length
 * 0 <= valueDiff <= 10^9
 *
 *
 */
// https://leetcode.cn/problems/contains-duplicate-iii/solutions/87257/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-46/
// @lc code=start
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	// /w得到桶id，但是 num<w时 num/w=0，此时会与正数/w=0的桶冲突，所以再-1，整体平移一位
	return x/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
