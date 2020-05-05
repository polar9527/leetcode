package main

func isMajorityElement(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mi := (lo + hi) / 2
		// [mi, hi] 中的元素都大于或者等于target, 缩小范围到[lo, mi-1]
		if target <= nums[mi] {
			// 临界状态
			// lo == hi , 此时 lo==hi==mi,
			// 如果target <= nums[mi] , 那么[：mi）都小于target, [mi:]都大于或者等于target
			hi = mi - 1
			// lo == mi， 结束循环
		}
		// [mi, hi] 中的元素都小于target, 缩小范围到[mi+1, hi]
		if target > nums[mi] {
			// 临界状态
			// lo == hi , 此时 lo==hi==mi,
			// 如果target > nums[mi] , 那么[:mi] 都小于target， [mi+1:]都大于target
			lo = mi + 1
			// lo == mi + 1， 结束循环
		}
	}
	if lo < len(nums) {
		// [lo:] 都大于或者等于target
		if nums[lo] == target {
			// first target index == lo
			if (lo + len(nums)/2) <= (len(nums) - 1) {
				if nums[lo+len(nums)/2] == target {
					return true
				}
			} else {
				// 	not majority
				return false
			}
		}
	}
	// target not exist
	return false
}
