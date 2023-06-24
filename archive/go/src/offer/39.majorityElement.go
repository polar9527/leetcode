/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (63.09%)
 * Likes:    582
 * Dislikes: 0
 * Total Accepted:    161.6K
 * Total Submissions: 256.1K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */
package main

import (
	"fmt"
	"math/rand"
)

// @lc code=start
func majorityElement(nums []int) int {
	return bit(nums)
	// return quickSelect(nums)
}

func bit(nums []int) int {

	cur := nums[0]
	count := 1
	for _, v := range nums[1:] {
		if count == 0 {
			cur = v
			count = 1
		} else if cur == v {
			count++
		} else {
			count--
		}
	}
	return cur
}

func quickSelect(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	lo := 0
	hi := len(nums) - 1
	n := len(nums)
	var pivot int
	for lo <= hi {
		pivot = partition(nums, lo, hi)
		// p == n/2 find return
		if pivot == n/2 {
			break
		}
		// p > p/2, cut
		if pivot > n/2 {
			hi = pivot - 1
		}
		// p < n/2, cut
		if pivot < n/2 {
			lo = pivot + 1
		}
	}
	return nums[pivot]
}

func partition(nums []int, lo, hi int) int {
	// select nums[rIndex] as pivot value
	// exchange it to head
	rIndex := rand.Intn(hi-lo+1) + lo
	pivotValue := nums[rIndex]
	nums[lo], nums[rIndex] = nums[rIndex], nums[lo]

	for lo < hi {
		for lo < hi && nums[hi] >= pivotValue {
			hi--
		}
		nums[lo] = nums[hi]

		for lo < hi && nums[lo] <= pivotValue {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = pivotValue
	return lo
}

// @lc code=end

func main() {

	ret := majorityElement([]int{3, 3, 4})
	fmt.Println(ret)
}
