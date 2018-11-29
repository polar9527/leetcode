package twosum

import "fmt"

// TwoSum solution of leetcode No.1 problem
func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	var result []int
	for index, num := range nums {
		fmt.Println(num)
		complement := target - num
		if _, ok := dict[complement]; !ok {
			dict[num] = index
			continue
		}
		result := []int{dict[complement], index}
		fmt.Println("in for")
		return result
	}
	fmt.Println("out for")
	return result
}
