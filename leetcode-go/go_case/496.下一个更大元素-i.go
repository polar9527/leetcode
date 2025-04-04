package go_case

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 *
 * https://leetcode.cn/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (71.81%)
 * Likes:    1060
 * Dislikes: 0
 * Total Accepted:    273.3K
 * Total Submissions: 380.5K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
 *
 * 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
 *
 * 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定
 * nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
 *
 * 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
 * 输出：[-1,3,-1]
 * 解释：nums1 中每个值的下一个更大元素如下所述：
 * - 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
 * - 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
 * - 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [2,4], nums2 = [1,2,3,4].
 * 输出：[3,-1]
 * 解释：nums1 中每个值的下一个更大元素如下所述：
 * - 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
 * - 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length <= nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 10^4
 * nums1和nums2中所有整数 互不相同
 * nums1 中的所有整数同样出现在 nums2 中
 *
 *
 *
 *
 * 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？
 *
 */

// @lc code=start
// func nextGreaterElement(nums1, nums2 []int) []int {
// 	// 维护一个栈，表示“待确定NGE (Next Greater Element) 的元素”，然后遍历序列。
// 	// 当我们碰上一个新元素，我们知道，越靠近栈顶的元素离新元素位置越近。
// 	// 所以不断比较新元素与栈顶，如果新元素比栈顶大，则可断定新元素就是栈顶的NGE，于是弹出栈顶并继续比较。
// 	// 直到新元素不比栈顶大，再将新元素压入栈。显然，这样形成的栈是单调递减的。

// 	mp := map[int]int{}
// 	stack := []int{}
// 	for i := len(nums2) - 1; i >= 0; i-- {
// 		num := nums2[i]
// 		// -------  A --------------
// 		// 不断拿栈顶元素与当前遍历到的元素做比较
// 		// 清除掉栈内存在的比当前遍历到的元素小的栈内元素
// 		// 直到栈为空，或者在栈内遇到比当前遍历到的元素 大的 栈内元素
// 		for len(stack) > 0 && num > stack[len(stack)-1] {
// 			// 此处 栈内元素全部是 单调递减的，栈底元素离当前所遍历到的元素nums[i]最远，栈顶元素最近
// 			stack = stack[:len(stack)-1]
// 		}
// 		// -------  A --------------

// 		// -------  B --------------
// 		// 根据题型，做不同处理
// 		// 此处 当清理完成栈内元素之后，当前的栈顶元素（如果有的话）就是 NGE (Next Greater Element)
// 		if len(stack) > 0 {
// 			mp[num] = stack[len(stack)-1]
// 		} else {
// 			mp[num] = -1
// 		}
// 		// -------  B --------------

// 		// -------  A --------------
// 		// 清理干净 栈内元素后，将当前遍历到的元素入栈

// 		stack = append(stack, num)
// 		// -------  A --------------
// 	}
// 	res := make([]int, len(nums1))
// 	for i, num := range nums1 {
// 		res[i] = mp[num]
// 	}
// 	return res
// }

// 反向遍历
// func nextGreaterElementR(nums1, nums2 []int) []int {
// 	// 维护一个栈，表示“待确定NGE (Next Greater Element) 的元素”，然后遍历序列。
// 	// 当我们碰上一个新元素，我们知道，越靠近栈顶的元素离新元素位置越近。
// 	// 所以不断比较新元素与栈顶，如果新元素比栈顶大，则可断定新元素就是栈顶的NGE，于是弹出栈顶并继续比较。
// 	// 直到新元素不比栈顶大，再将新元素压入栈。显然，这样形成的栈是单调递减的。

// 	mp := map[int]int{}
// 	// for _, num := range nums2 {
// 	// 	mp[num] = -1
// 	// }
// 	stack := []int{}
// 	for i := 0; i < len(nums2); i++ {
// 		num := nums2[i]
// 		// -------  A --------------
// 		// 不断拿栈顶元素与当前遍历到的元素做比较
// 		// 清除掉栈内存在的比当前遍历到的元素小的栈内元素
// 		// 直到栈为空，或者在栈内遇到比当前遍历到的元素 大的 栈内元素
// 		for len(stack) > 0 && num > stack[len(stack)-1] {
// 			// -------  B --------------
// 			// 根据题型，做不同处理
// 			// 此处 栈内元素全部是 单调递减的，栈底元素离当前所遍历到的元素nums[i]最远，栈顶元素最近
// 			// 也就是说存储的都是已经遍历过的，并且都还没有找到自己右侧比自己大的第一个元素NGE
// 			// 记录当前 遍历到的元素 是当前栈顶元素的右侧的第一个最大元素NGE (Next Greater Element)
// 			mp[stack[len(stack)-1]] = num
// 			// -------  B --------------

// 			stack = stack[:len(stack)-1]
// 		}
// 		// -------  A --------------
// 		// -------  A --------------
// 		// 清理干净 栈内元素后，将当前遍历到的元素入栈

// 		stack = append(stack, num)
// 		// -------  A --------------
// 	}
// 	res := make([]int, len(nums1))
// 	for i, num := range nums1 {
// 		if v, ok := mp[num]; ok {
// 			res[i] = v
// 		} else {
// 			res[i] = -1
// 		}

// 	}
// 	return res
// }

func nextGreaterElement(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	set1 := make(map[int]int)

	for i := range nums1 {
		set1[nums1[i]] = i
	}
	// 通过 set2[nums1[i]] 定位 nums2 中的index

	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(nums2); i++ {
		if nums2[stack[len(stack)-1]] >= nums2[i] {
			stack = append(stack, i)
		} else {
			// 栈顶存储的 nums2 中的 index = stack[len(stack)-1]
			// 当前的 nums2[i] 是 nums2[stack[len(stack)-1]] 的右边首个最大值
			for len(stack) != 0 && nums2[stack[len(stack)-1]] < nums2[i] {
				// collect result
				// 找到 nums1 中对应的 index
				if index, ok := set1[nums2[stack[len(stack)-1]]]; ok {
					res[index] = nums2[i]
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return res
}

// @lc code=end
