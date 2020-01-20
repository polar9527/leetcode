/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
// package main

// import "fmt"

// @lc code=start
func isPalindrome(x int) bool {

	if x >= 0 && x < 10 {
		return true
	}

	if x < 0 || x == 10 {
		return false
	}

	numberArray := []int{}
	for iter := x; iter > 0; {
		// fmt.Println(iter)
		numberArray = append(numberArray, iter%10)
		// fmt.Println(numberArray)
		iter = iter / 10
	}

	arrayLenght := len(numberArray)
	for head, tail := 0, arrayLenght-1; head < tail; {
		if numberArray[head] != numberArray[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// func main() {
// 	number1 := 12321
// 	number2 := 123321
// 	fmt.Println(isPalindrome(number1))
// 	fmt.Println(isPalindrome(number2))
// 	number1 = 123421
// 	number2 = -12333421
// 	fmt.Println(isPalindrome(number1))
// 	fmt.Println(isPalindrome(number2))
// }

// @lc code=end
