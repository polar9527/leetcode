/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {

	return singleNumber2(nums)
}

func singleNumber1(nums []int) int {

	once, twice := 0, 0

	for _, v := range nums {
		// #1 初始状态，once 和 twice 这时候管理的位都是0
		// #1 once ^ v 该位置1
		// #1 ~twice & (once ^ v)， ~twice这位置1，目的是通过 & 放过这位给到once
		// #2 此时once这位为1， once ^ v 将这位置0
		// #2 此时 twice 这位为0，~twice 通过 & 将这位放多，最终得到once这位为0
		// #3 此时 once 这位置为0， once ^ v 该位置1
		// #3 此时 twice 这位置为1，~twice 通过& 将放过给到once,即 once 这位为 1
		once = ^twice & (once ^ v)
		// #1 此时 twice ^ v 该位置1
		// #1 ~once 这位置为0， 目的是通过&将这位在twice 中置0
		// #2 此时twice这位为0，twice ^ v 将这位置1
		// #2 此时once这位在上一步变为0，~once 通过& 将这位放过给到twice,即twice中这位为1
		// #3 此时twice 这位为1， twice ^ v 将这位置0
		// #3 此时once这位在上一步变为1，~once 通过& 将这位放过给到twice，即twice中这位为0
		twice = ^once & (twice ^ v)
		// #1 此时 once 和 twice 这位分别是1和0
		// #2 此时 once 和 twice 这位分别是0和1
		// #2 此时 once 和 twice 这位分别是1和0
	}
	return once
}

func singleNumber2(nums []int) int {
	set := make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}
	s
	sumOfSet := 0
	for _, v := range set {
		sumOfSet += v
	}

	sumOfNums := 0
	for k, _ := range nums {
		sumOfNums += k
	}
	return (3*sumOfSet - sumOfNums) / 2
}

func singleNumber2(nums []int) int {
	return
}

// @lc code=end

