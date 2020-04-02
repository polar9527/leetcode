package main

func exchange(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	for lo, hi := 0, l-1; lo < hi; {
		for lo < hi && nums[lo]%2 != 0 {
			lo++
		}
		for lo < hi && nums[hi]%2 == 0 {
			hi--
		}

		if lo < hi && nums[lo]%2 == 0 && nums[hi]%2 != 0 {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			lo++
			hi--
		}

	}
	return nums
}

// func main() {
// 	arr := []int{1,2,3,4}
//
// 	exchange(arr)
// }
