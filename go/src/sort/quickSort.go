package sort

import (
	"math/rand"
	"time"
)

type Partition func([]int) int

func quickSort(nums []int, pFunc Partition) {
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
