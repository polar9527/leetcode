package sort

import (
	"math/rand"
	"time"
)

type partition func([]int) int

func quickSort(nums []int, pFunc partition) {
	length := len(nums)
	if length < 2 {
		return
	}

	pivot := pFunc(nums)

	quickSort(nums[0:pivot], pFunc)
	quickSort(nums[pivot+1:], pFunc)
}

// lazy switch
func partitionLazySwitch(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())

	head := 0
	tail := l - 1

	pivotIndex := rand.Intn(l)
	pivot := nums[pivotIndex]
	nums[pivotIndex] = nums[head]

	for head < tail {
		for head < tail && nums[tail] >= pivot {
			tail--
		}
		// 1. head == tail
		// 2. head < tail && nums[tail] < pivot
		nums[head] = nums[tail]
		if head < tail {
			head++
		}
		for head < tail && nums[head] <= pivot {
			head++
		}
		nums[tail] = nums[head]
		if head < tail {
			tail--
		}
	}
	nums[head] = pivot
	return head
}

// Aggressive switch
// 应对全部是重复元素的最坏情况，这种情况下 partitionL 可能一直将nums 切分为 1：len(nums)-1，算法复杂度会退化为O(n2)
// 而partitionA 可以保证 nums被对半切分，从而保证这种情况下算法复杂度仍然为O(nlogn), 但是会加剧快排算法的不稳定性
func partitionAggressiveSwitch(nums []int) int {
	// len(nums) >= 3
	length := len(nums)
	if length <= 1 {
		return 0
	}

	head := 0
	tail := len(nums) - 1

	rand.Seed(time.Now().UnixNano())
	rIndex := rand.Intn(len(nums))
	pivot := nums[rIndex]
	nums[rIndex] = nums[head]

	for head < tail {
		for head < tail && nums[tail] > pivot {
			tail--
		}
		// 1. head == tail
		// 2. head < tail && nums[tail] <= pivot
		nums[head] = nums[tail]
		if head < tail {
			head++
		}
		for head < tail && nums[head] < pivot {
			head++
		}
		nums[tail] = nums[head]
		if head < tail {
			tail--
		}
	}
	nums[head] = pivot
	return head
}

// func main() {
// 	// arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
// 	arr := []int{1, 1}
// 	containsDuplicateSort(arr)
// 	fmt.Println(arr)
//
// }
