package sort

import (
	"math/rand"
	"time"
)

type partition func([]int) int

func quickSort(nums []int, pFunc partition) {
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

	pivot := pFunc(nums)

	quickSort(nums[0:pivot], pFunc)
	quickSort(nums[pivot+1:], pFunc)

}

// lazy switch
func partitionL(nums []int) int {
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

	rand.Seed(time.Now().UnixNano())
	rIndex := rand.Intn(len(nums))
	pivot := nums[rIndex]
	nums[0], nums[rIndex] = nums[rIndex], nums[0]

	for head < tail {
		for head < tail && nums[tail] >= pivot {
			tail--
		}
		nums[head] = nums[tail]
		for head < tail && nums[head] <= pivot {
			head++
		}
		nums[tail] = nums[head]
	}
	nums[head] = pivot
	return head
}

// Aggressive switch
// 应对全部是重复元素的最坏情况，这种情况下 partitionL 可能一直将nums 切分为 1：len(nums)-1，算法复杂度会退化为O(n2)
// 而partitionA 可以保证 nums被对半切分，从而保证这种情况下算法复杂度仍然为O(nlogn)
func partitionA(nums []int) int {
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

	rand.Seed(time.Now().UnixNano())
	rIndex := rand.Intn(len(nums))
	pivot := nums[rIndex]
	nums[0], nums[rIndex] = nums[rIndex], nums[0]

	for head < tail {
		for head < tail {
			if nums[tail] > pivot {
				tail--
			} else {
				nums[head] = nums[tail]
				head++
				break
			}
		}

		for head < tail {
			if nums[head] < pivot {
				head++
			} else {
				nums[tail] = nums[head]
				tail--
				break
			}
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
