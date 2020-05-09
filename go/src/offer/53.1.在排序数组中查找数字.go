package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	first := getFirst(nums, target)
	if first < 0 {
		return 0
	}
	last := getLast(nums, target)
	if last > len(nums)-1 {
		return 0
	}
	result := last - first + 1
	return result
}

func getFirst(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mi := (lo + hi) / 2
		if target > nums[mi] {
			// 	左边所有的都要小于target
			lo = mi + 1
		} else {
			// target <= nums[mi]
			hi = mi - 1
		}
		// 	当 start == end 时 mi = start = end, [:lo)都小于target , [hi:]都大于等于target
	}
	return lo
}

func getLast(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mi := (lo + hi) / 2
		if target < nums[mi] {
			hi = mi - 1
		} else {
			// target >= nums[mi]
			lo = mi + 1
		}
	}
	return hi
}

func main() {
	// search([]int{2,2}, 3)
	search([]int{4, 4}, 3)
}
