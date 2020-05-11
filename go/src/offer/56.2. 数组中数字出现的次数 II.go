package main

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	bits := make([]int, 32)
	for _, num := range nums {
		mask := 1
		for j := 0; j < 32; j++ {
			if num&mask != 0 {
				bits[j] += 1
			}
			mask <<= 1
		}
	}
	number := 0
	for i := 31; i >= 0; i-- {

		if bits[i]%3 != 0 {
			number += 1
		}
		if i == 0 {
			break
		}
		number <<= 1
	}
	return number
}

func main() {
	singleNumber([]int{3, 4, 3, 3})
}
