package go_case

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
 *
 * algorithms
 * Easy (56.42%)
 * Likes:    3575
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 3.3M
 * Testcase Example:  '[1,1,2]'
 *
 * 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序
 * 应该保持 一致 。然后返回 nums 中唯一元素的个数。
 *
 * 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
 *
 *
 * 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums
 * 的大小不重要。
 * 返回 k 。
 *
 *
 * 判题标准:
 *
 * 系统会用下面的代码来测试你的题解:
 *
 *
 * int[] nums = [...]; // 输入数组
 * int[] expectedNums = [...]; // 长度正确的期望答案
 *
 * int k = removeDuplicates(nums); // 调用
 *
 * assert k == expectedNums.length;
 * for (int i = 0; i < k; i++) {
 * ⁠   assert nums[i] == expectedNums[i];
 * }
 *
 * 如果所有断言都通过，那么您的题解将被 通过。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：2, nums = [1,2,_]
 * 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,1,1,1,2,2,3,3,4]
 * 输出：5, nums = [0,1,2,3,4]
 * 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4
 * 。不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 已按 非严格递增 排列
 *
 *
 */

// @lc code=start
// func removeDuplicates(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	slow := 0
// 	val := nums[0]
// 	for fast := 1; fast < len(nums); fast++ {
// 		if nums[fast] > val {
// 			slow++
// 			nums[slow] = nums[fast]
// 			val = nums[fast]
// 		}
// 	}
// 	return slow + 1
// }

func removeDuplicates(nums []int) int {

	l := len(nums)
	if l == 0 {
		return 0
	}

	// 当 l > 0 的时候，最后的数组里至少剩下1个数字
	// fast 指向 当前遍历到的数字，
	// slow 指向 下一个不同数字将要填入的位置
	slow, fast := 1, 1
	for fast < l {
		// 因为 数组中可能有相同的数字，所以fast 至少不会 比 slow 跑得慢
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end
