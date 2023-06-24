package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	numbers := make([]int, 0)
	if n <= 0 {
		return numbers
	}
	numberBytes := make([]byte, n)

	for i := range numberBytes {
		numberBytes[i] = '0'
	}

	for !increment(numberBytes) {
		number := getNumber(numberBytes)
		numbers = append(numbers, number)
	}

	return numbers
}

func increment(numberBytes []byte) bool {
	length := len(numberBytes)
	isOverFlow := false
	carry := 0
	for index := length - 1; index >= 0; index-- {

		// 处于某个十进制位时的step动作
		// 其中个位的step动作和其他位的step动作是有区别的
		sum := int(numberBytes[index]-'0') + carry
		if index == length-1 {
			sum++
		}

		// 进位时的操作
		if sum >= 10 {
			// 最高位的时候，溢出
			if index == 0 {
				isOverFlow = true
				// 其他位做进位处理，并将sum重置
			} else {
				carry = 1
				sum -= 10
				// 更新数字
				numberBytes[index] = byte(sum) + '0'
			}
		} else {
			// 不进位的时候，更新数字
			numberBytes[index] = byte(sum) + '0'
			break
		}

	}

	return isOverFlow
}

func getNumber(numberBytes []byte) int {
	length := len(numberBytes)
	beginning := 0
	// 找到最高位
	for i := 0; i < length; i++ {
		if numberBytes[i] != '0' {
			beginning = i
			break
		}
	}

	// 从个位开始计算数值
	base := 1
	sum := 0
	for i := length - 1; i >= beginning; i-- {
		sum += int(numberBytes[i]-'0') * base
		base *= 10
	}

	return sum
}

func printNumbersR(n int) []int {
	numbers := make([]int, 0)
	numberBytes := make([]byte, n)

	for i := range numberBytes {
		numberBytes[i] = '0'
	}

	if n <= 0 {
		return numbers
	}

	for i := 0; i < 10; i++ {
		numberBytes[0] = byte(i) + '0'
		printNumbersRecursive(numberBytes, n, 1, &numbers)
	}

	return numbers[1:]
}

func printNumbersRecursive(numberBytes []byte, length, index int, numbers *[]int) {
	if index == length {
		number := getNumber(numberBytes)
		*numbers = append(*numbers, number)
		return
	}

	for i := 0; i < 10; i++ {
		numberBytes[index] = byte(i) + '0'
		printNumbersRecursive(numberBytes, length, index+1, numbers)
	}
}

func main() {
	arr := printNumbers(2)
	fmt.Println(arr)
	arrR := printNumbersR(2)
	fmt.Println(arrR)

}
