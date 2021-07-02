/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (54.44%)
 * Likes:    639
 * Dislikes: 649
 * Total Accepted:    474.4K
 * Total Submissions: 864.4K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 *
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 */
package leetcode

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// @lc code=start
func containsDuplicate(nums []int) bool {
	return containsDuplicateHashMap(nums)
}

func containsDuplicateSort(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	sort.Sort(sort.IntSlice(nums))
	// quickSort(nums)
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func quickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	if length == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}

	pivot := partition(nums)

	quickSort(nums[0:pivot])
	quickSort(nums[pivot+1:])

}

func partition(nums []int) int {
	// len(nums) >= 3
	length := len(nums)
	if length <= 1 {
		return 0
	}
	if length == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return 0
	}
	head := 0
	tail := len(nums) - 1

	rand.Seed(time.Now().Unix())
	rIndex := rand.Intn(len(nums))
	pivot := nums[rIndex]
	nums[0], nums[rIndex] = nums[rIndex], nums[0]

	for head < tail {
		for head < tail && nums[tail] > pivot {
			tail--
		}
		nums[head] = nums[tail]
		head++
		for head < tail && nums[head] < pivot {
			head++
		}
		nums[tail] = nums[head]
		tail--
	}
	nums[head] = pivot
	return head
}

func containsDuplicateHashMap(nums []int) bool {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}

	return false
}

// @lc code=end

func main() {
	// arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	arr := []int{1, 1}
	containsDuplicateSort(arr)
	fmt.Println(arr)

}
